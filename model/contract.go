package model

import "time"

// GetTradeRecords account's trade records
func GetTradeRecordsByContract(contract string, count int) ([]Trade, error) {
	rows, err := db.Query("select tx,direction,amount,price,block,ts from trade where contract=? order by ts desc limit ?", contract, count)
	if err != nil {
		return nil, err
	}
	data := make([]Trade, 0, 1)
	for rows.Next() {
		t := Trade{}
		rows.Scan(&t.Tx, &t.Direction, &t.Amount, &t.Price, &t.Block, &t.Ts)
		data = append(data, t)
	}
	return data, nil
}

func GetExplosiveRecordsByContract(contract string, count int) ([]Trade, error) {
	rows, err := db.Query("select direction,amount,price,block,ts from explosive where contract=? order by ts desc limit ?", contract, count)
	if err != nil {
		return nil, err
	}
	data := make([]Trade, 0, 1)
	for rows.Next() {
		t := Trade{}
		rows.Scan(&t.Direction, &t.Amount, &t.Price, &t.Block, &t.Ts)
		data = append(data, t)
	}
	return data, nil
}

func GetLatestContractUpdateTime() (time.Time, error) {
	row := db.QueryRow("select ts from contract")
	var s string
	if err := row.Scan(&s); err != nil {
		return time.Now(), err
	}
	return time.Parse("2006-01-02 15:04:05", s)
}
