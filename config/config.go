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

type hedgex struct {
	Address    string `json:"address"`
	MarginCoin string `json:"margin_coin"`
	TradeCoin  string `json:"trade_coin"`
}

type chainNode struct {
	Https string `json:"http"`
	Wss   string `json:"ws"`
}

type explosive struct {
	Tick      time.Duration `json:"tick"`
	ToAddress string        `json:"to_address"`
}

type interest struct {
	Tick      time.Duration `json:"tick"`
	Begin     int64         `json:"begin"`
	End       int64         `json:"end"`
	ToAddress string        `json:"to_address"`
}

type testcoin struct {
	Count       int    `json:"count"`
	CoinAnount  string `json:"coin_amount"`
	Token       string `json:"token"`
	TokenAmount string `json:"token_amount"`
}

var (
	//Service, its'value indicate running different service.
	//1 is only run the public service that not need privatekey of wallet;
	//2 only run explosive and interest service that need privatekey of wallet;
	//3 run both 1 and 2;
	Service             int
	Db                  db
	HttpPort            int
	WsPort              int
	WsTick              time.Duration
	Explosive           explosive
	Interest            interest
	MaxKlineCount       int
	MaxTradeRecordCount int
	ChainNode           chainNode
	Contract            []hedgex
	PrivateKey          string
	TestCoin            testcoin
)

//Load load config file
func Load() {
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
		Explosive           explosive     `json:"explosive"`
		Interest            interest      `json:"interest"`
		ExplosiveTo         string        `json:"explosive_to_address"`
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
	Explosive = all.Explosive
	Interest = all.Interest
	ChainNode = all.ChainNode
	Contract = all.Contract
	PrivateKey = all.PrivateKey
	TestCoin = all.TestCoin
	if MaxKlineCount < 1 {
		log.Panic("max kline count must > 0")
	}
}
