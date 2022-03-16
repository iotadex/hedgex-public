package model

import (
	"database/sql"
	"fmt"
	"hedgex-public/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectToMysql() {
	var err error
	db, err = sql.Open("mysql", config.Db.Usr+":"+config.Db.Pwd+"@tcp("+config.Db.Host+":"+config.Db.Port+")/"+config.Db.DbName)
	if err != nil {
		log.Panic(err)
	}
	db.SetMaxOpenConns(config.Db.OpenConns)
	db.SetMaxIdleConns(config.Db.IdleConns)
	db.SetConnMaxLifetime(config.Db.LifeTime * time.Second)

	if err = db.Ping(); nil != err {
		log.Panic("Connect to Mysql error : " + err.Error())
	}
}

func Ping() error {
	if db == nil {
		return fmt.Errorf("mysql connection is nil")
	}
	if err := db.Ping(); nil != err {
		return fmt.Errorf("connect to Mysql error : %v", err)
	}
	return nil
}
