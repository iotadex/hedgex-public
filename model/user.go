package model

import (
	"database/sql"
	"errors"
)

/*
CREATE TABLE user (
  `account` varchar(45) NOT NULL COMMENT '用户地址，格式为：0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e',
  `contract` varchar(45) NOT NULL COMMENT '所属合约',
  `margin` bigint NOT NULL DEFAULT '0' COMMENT '保证金数量',
  `lposition` int NOT NULL DEFAULT '0' COMMENT '多仓持仓量',
  `lprice` bigint NOT NULL DEFAULT '0' COMMENT '多仓持仓价',
  `sposition` int NOT NULL DEFAULT '0' COMMENT '空仓持仓量',
  `sprice` bigint NOT NULL DEFAULT '0' COMMENT '空仓持仓价',
  `block` bigint NOT NULL DEFAULT '0' COMMENT '最后一次更新的区块高度',
  PRIMARY KEY (`account`,`contract`),
  KEY `block` (`block)
)
*/

type PositionStat struct {
	Long  int `json:"long"`
	Short int `json:"short"`
	Total int `json:"total"`
}

func GetStatPositions() (map[string]*PositionStat, error) {
	rows, err := db.Query("select contract,count(account) as lc from user where lposition>0 group by contract")
	if err != nil {
		return nil, err
	}
	data := make(map[string]*PositionStat)
	for rows.Next() {
		var contract string
		var lc int
		err := rows.Scan(&contract, &lc)
		if err != nil {
			return nil, err
		}
		if _, exist := data[contract]; !exist {
			data[contract] = &PositionStat{}
		}
		data[contract].Long = lc
	}

	rows, err = db.Query("select contract,count(account) as lc from user where sposition>0 group by contract")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var contract string
		var lc int
		err := rows.Scan(&contract, &lc)
		if err != nil {
			return nil, err
		}
		if _, exist := data[contract]; !exist {
			data[contract] = &PositionStat{}
		}
		data[contract].Short = lc
	}

	rows, err = db.Query("select contract,count(account) as lc from user where lposition>0 or sposition>0 group by contract")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var contract string
		var lc int
		err := rows.Scan(&contract, &lc)
		if err != nil {
			return nil, err
		}
		if _, exist := data[contract]; !exist {
			data[contract] = &PositionStat{}
		}
		data[contract].Total = lc
	}

	return data, nil
}

//GetLastBlock get the last block from use table
func GetLastBlock(contract string) (int64, error) {
	row := db.QueryRow("select block from user order by block desc limit 1")
	var maxBlock int64
	err := row.Scan(&maxBlock)
	if err == sql.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return maxBlock, nil
}

type User struct {
	Account     string
	Margin      int64
	Lposition   uint64
	Lprice      uint64
	Sposition   uint64
	Sprice      uint64
	InterestDay uint
	Block       uint64
}

//GetUsers get users
func GetUsers(contract string) ([]User, uint64, error) {
	rows, err := db.Query("SELECT account,margin,lposition,lprice,sposition,sprice,interest_day,block FROM user where contract=" + "'" + contract + "'")
	if err != nil {
		return nil, 0, err
	}
	data := make([]User, 0)
	var maxBlock uint64
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Account, &u.Margin, &u.Lposition, &u.Lprice, &u.Sposition, &u.Sprice, &u.InterestDay, &u.Block)
		if err != nil {
			return nil, 0, err
		}
		data = append(data, u)
		if u.Block > maxBlock {
			maxBlock = u.Block
		}
	}
	return data, maxBlock, nil
}

//UpdateUser update user's data
func UpdateUser(contract string, u *User) error {
	_, err := db.Exec("replace into user(account,contract,margin,lposition,lprice,sposition,sprice,interest_day,block) values(?,?,?,?,?,?,?,?,?) ",
		u.Account, contract, u.Margin, u.Lposition, u.Lprice, u.Sposition, u.Sprice, u.InterestDay, u.Block)
	return err
}

//InsertRecharge insert a recharge record
func InsertRecharge(tx string, contract string, account string, amount uint64, block uint64) error {
	_, err := db.Exec("insert into recharge(tx,contract,account,amount,block) values(?,?,?,?,?)", tx, contract, account, amount, block)
	return err
}

//InsertWithdraw insert a withdraw record
func InsertWithdraw(tx string, contract string, account string, amount uint64, block uint64) error {
	_, err := db.Exec("insert into withdraw(tx,contract,account,amount,block) values(?,?,?,?,?)", tx, contract, account, amount, block)
	return err
}

// InsertTrade insert a trade record
func InsertTrade(tx string, contract string, account string, direction int8, amount uint64, price uint64, block uint64) error {
	_, err := db.Exec("insert into trade(tx,contract,account,direction,amount,price,block) values(?,?,?,?,?,?,?)", tx, contract, account, direction, amount, price, block)
	return err
}

type Trade struct {
	Tx        string `json:"tx"`
	Contract  string `json:"contract"`
	Account   string `json:"account"`
	Direction int8   `json:"direction"`
	Amount    uint64 `json:"amount"`
	Price     uint64 `json:"price"`
	Block     uint64 `json:"block"`
}

// GetTradeRecords account's trade records
func GetTradeRecords(contract string, account string, count int) ([]Trade, error) {
	rows, err := db.Query("select tx,direction,amount,price,block from trade where contract=? and account=? order by block desc limit ?", contract, account, count)
	if err != nil {
		return nil, err
	}
	data := make([]Trade, 0, 1)
	for rows.Next() {
		t := Trade{}
		rows.Scan(&t.Tx, &t.Direction, &t.Amount, &t.Price, &t.Block)
		data = append(data, t)
	}
	return data, nil
}

// InsertTrade insert a trade record
func InsertExplosive(tx string, contract string, account string, direction int8, amount uint64, price uint64, block uint64) error {
	_, err := db.Exec("insert into explosive(tx,contract,account,direction,amount,price,block) values(?,?,?,?,?,?,?)", tx, contract, account, direction, amount, price, block)
	return err
}

func GetExplosive(contract string, account string, count int) ([]Interest, error) {
	rows, err := db.Query("select tx,direction,amount,price,block from explosive where contract=? and account=? order by block desc limit ?", contract, account, count)
	if err != nil {
		return nil, err
	}
	data := make([]Interest, 0, 1)
	for rows.Next() {
		t := Interest{}
		rows.Scan(&t.Tx, &t.Direction, &t.Amount, &t.Price, &t.Block)
		data = append(data, t)
	}
	return data, nil
}

// InsertTrade insert a trade record
func InsertInterest(tx string, contract string, account string, direction int8, amount uint64, price uint64, block uint64) error {
	_, err := db.Exec("insert into interest(tx,contract,account,direction,amount,price,block) values(?,?,?,?,?,?,?)", tx, contract, account, direction, amount, price, block)
	return err
}

type Interest struct {
	Tx        string `json:"tx"`
	Direction int8   `json:"direction"`
	Amount    uint64 `json:"amount"`
	Price     uint64 `json:"price"`
	Block     uint64 `json:"block"`
}

func GetInterests(contract string, account string, count int) ([]Interest, error) {
	rows, err := db.Query("select tx,direction,amount,price,block from interest where contract=? and account=? order by block desc limit ?", contract, account, count)
	if err != nil {
		return nil, err
	}
	data := make([]Interest, 0, 1)
	for rows.Next() {
		t := Interest{}
		rows.Scan(&t.Tx, &t.Direction, &t.Amount, &t.Price, &t.Block)
		data = append(data, t)
	}
	return data, nil
}

// UpdateTestCoin update testcoin send count
func UpdateTestCoin(account string) error {
	row, err := db.Query("select count from testcoin where account=?", account)
	if err != nil {
		return err
	}
	var count int
	if row.Next() {
		if err = row.Scan(&count); err != nil {
			return err
		}
	}
	if count >= 3 {
		return errors.New("over the number for getting test coins")
	}
	count++
	_, err = db.Exec("replace into testcoin(account,count) values(?,?)", account, count)
	return err
}
