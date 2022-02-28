package model

// GetIpCount get ip's count in the table email
func GetIpCount(ip string) (int, error) {
	count := 0
	row := db.QueryRow("select count(addr) from email where ip=?", ip)
	err := row.Scan(&count)
	return count, err
}

// InsertEmail insert a email address to the table
func InsertEmail(email, ip string) error {
	_, err := db.Exec("replace into email(addr,ip) values(?,?)", email, ip)
	return err
}

// GetEmails, get email between from and to by ts
func GetEmails(from, to string) ([]string, error) {
	rows, err := db.Query("select addr from email where ts>=? and ts<?", from, to)
	if err != nil {
		return nil, err
	}
	emails := make([]string, 0)
	for rows.Next() {
		var addr string
		if err := rows.Scan(&addr); err != nil {
			return nil, err
		}
		emails = append(emails, addr)
	}
	return emails, nil
}

func GetAccountTestCoinSendCount(account string) (int, error) {
	count := 0
	row := db.QueryRow("select count from testcoin where account=?", account)
	err := row.Scan(&count)
	return count, err
}

func IncreaseTestCoinCount(account string) error {
	_, err := db.Exec("update testcoin set count=count-1 where account=?", account)
	return err
}