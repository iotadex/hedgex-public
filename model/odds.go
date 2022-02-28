package model

import "database/sql"

// GetIpCount get ip's count in the table email
func GetIpCount(ip string) (int, error) {
	count := 0
	row := db.QueryRow("select count(addr) from email where ip=?", ip)
	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
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
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return count, err
}

func IncreaseTestCoinCount(account string, count int) error {
	_, err := db.Exec("replace into testcoin(account,count) values(?,?)", account, count)
	return err
}
