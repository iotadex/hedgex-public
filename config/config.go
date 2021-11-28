package config

import (
	"encoding/json"
	"hedgex-server/tools"
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

type hedgex struct {
	Address    string `json:"address"`
	MarginCoin string `json:"margin_coin"`
	TradeCoin  string `jons:"trade_coin"`
}

type chainNode struct {
	Https string `json:"http"`
	Wss   string `json:"ws"`
}

type testcoin struct {
	Count       int    `json:"count"`
	CoinAnount  string `json:"coin_amount"`
	Token       string `json:"token"`
	TokenAmount string `json:"token_amount"`
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
	ChainNode           chainNode
	Contract            []hedgex
	PrivateKey          string
	TestCoin            testcoin
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
		ChainNode           chainNode     `json:"chain_node"`
		Contract            []hedgex      `json:"contract"`
		PrivateKey          string        `json:"wallet"`
		TestCoin            testcoin      `json:"testcoin"`
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
	TestCoin = all.TestCoin
	if MaxKlineCount < 1 {
		log.Panic("max kline count must > 0")
	}

	//if you have yourself key, you can use it.
	//the package of "tools" is private.
	key := tools.InputKey()
	PrivateKey = tools.AesCBCDecrypt(PrivateKey, key)
}
