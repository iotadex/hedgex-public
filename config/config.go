package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

//db database config
type db struct {
	LifeTime  time.Duration `json:"lifetime"`
	OpenConns int           `json:"openconns"`
	IdleConns int           `json:"idleconns"`
	Host      string        `json:"host"`
	Port      string        `json:"port"`
	DbName    string        `json:"dbname"`
	Usr       string        `json:"usr"`
	Pwd       string        `json:"pwd"`
}

type contract struct {
	Pair  []string `json:"pair"`
	Https string   `json:"http"`
	Wss   string   `json:"ws"`
}

var (
	Service             int
	Db                  db
	HttpPort            int
	WsPort              int
	WsTick              time.Duration
	ExplosiveTick       time.Duration
	MaxKlineCount       int
	MaxTradeRecordCount int
	Contract            contract
	PrivateKey          string
)

// LoadConfig load config file
func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	type Config struct {
		Service             int           `json:"service"`
		HttpPort            int           `json:"http_port"`
		WsPort              int           `json:"ws_port"`
		WsTick              time.Duration `json:"ws_tick"`
		ExplosiveTick       time.Duration `json:"explosive_tick"`
		KlineMaxCount       int           `json:"kline_max_count"`
		MaxTradeRecordCount int           `json:"max_trade_count"`
		Db                  db            `json:"db"`
		Contract            contract      `json:"contract"`
		PrivateKey          string        `json:"wallet"`
	}
	all := &Config{}
	if err = json.NewDecoder(file).Decode(all); err != nil {
		log.Panic(err)
	}
	Service = all.Service
	Db = all.Db
	HttpPort = all.HttpPort
	WsPort = all.WsPort
	WsTick = all.WsTick
	MaxKlineCount = all.KlineMaxCount
	Contract = all.Contract
	PrivateKey = all.PrivateKey
	if MaxKlineCount < 1 {
		log.Panic("max kline count must > 0")
	}
}