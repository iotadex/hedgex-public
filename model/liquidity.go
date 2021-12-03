package model

func InsertMint(tx string, contract string, account string, amount uint64, block uint64) error {
	_, err := db.Exec("insert into mint(tx,contract,account,amount,block) values(?,?,?,?,?)", tx, contract, account, amount, block)
	return err
}

func InsertBurn(tx string, contract string, account string, amount uint64, block uint64) error {
	_, err := db.Exec("insert into burn(tx,contract,account,amount,block) values(?,?,?,?,?)", tx, contract, account, amount, block)
	return err
}
