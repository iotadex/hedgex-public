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
	Name  string `json:"name"`
	Https string `json:"http"`
	Wss   string `json:"ws"`
}

type testcoin struct {
	Count       int    `json:"count"`
	CoinAmount  string `json:"coin_amount"`
	Token       string `json:"token"`
	TokenAmount string `json:"token_amount"`
}

var (
	Env           string
	Db            db
	HttpPort      int
	WsPort        int
	WsTick        time.Duration
	MaxKlineCount int
	IndexTick     time.Duration
	ChainNode     chainNode
	Contract      []hedgex
	PrivateKey    string
	TestCoin      testcoin
)

//Load load config file
func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	type Config struct {
		Env           string        `json:"env"`
		Service       int           `json:"service"`
		HttpPort      int           `json:"http_port"`
		WsPort        int           `json:"ws_port"`
		WsTick        time.Duration `json:"ws_tick"`
		ExplosiveTo   string        `json:"explosive_to_address"`
		KlineMaxCount int           `json:"kline_max_count"`
		IndexTick     time.Duration `json:"index_tick"`
		Db            db            `json:"db"`
		ChainNode     chainNode     `json:"chain_node"`
		Contract      []hedgex      `json:"contract"`
		PrivateKey    string        `json:"wallet"`
		TestCoin      testcoin      `json:"testcoin"`
	}
	all := &Config{}
	if err = json.NewDecoder(file).Decode(all); err != nil {
		log.Panic(err)
	}
	Env = all.Env
	Db = all.Db
	HttpPort = all.HttpPort
	WsPort = all.WsPort
	WsTick = all.WsTick
	MaxKlineCount = all.KlineMaxCount
	IndexTick = all.IndexTick
	ChainNode = all.ChainNode
	Contract = all.Contract
	PrivateKey = all.PrivateKey
	TestCoin = all.TestCoin
	if MaxKlineCount < 1 {
		log.Panic("max kline count must > 0")
	}
}
