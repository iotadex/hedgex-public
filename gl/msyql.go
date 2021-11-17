package gl

import (
	"database/sql"
	"hedgex-server/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectToMysql 连接mysql
func init() {
	// 创建数据库连接
	var err error
	DB, err = sql.Open("mysql", config.Db.Usr+":"+config.Db.Pwd+"@tcp("+config.Db.Host+":"+config.Db.Port+")/"+config.Db.DbName)
	if err != nil {
		log.Panic(err)
	}
	// 最大连接数
	DB.SetMaxOpenConns(config.Db.OpenConns)
	// 闲置连接数
	DB.SetMaxIdleConns(config.Db.IdleConns)
	// 最大连接周期
	DB.SetConnMaxLifetime(config.Db.LifeTime * time.Second)

	if err = DB.Ping(); nil != err {
		log.Panic("Connect to Mysql error : " + err.Error())
	}
}
