package config

import (
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
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

type test struct {
	LimitCount    int    `json:"limit_count"`
	SendAmount    int64  `json:"send_amount"`
	Token         string `json:"token"`
	Wallet        string `json:"wallet"`
	PrivateKey    *ecdsa.PrivateKey
	PublicAddress common.Address
}

var (
	Env           string
	Db            db
	HttpPort      int
	WsPort        int
	WsTick        time.Duration
	MaxKlineCount int
	ChainNode     string
	Contract      []hedgex
	IpLimit       int
	Test          test
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
		HttpPort      int           `json:"http_port"`
		WsPort        int           `json:"ws_port"`
		WsTick        time.Duration `json:"ws_tick"`
		KlineMaxCount int           `json:"kline_max_count"`
		Db            db            `json:"db"`
		ChainNode     string        `json:"chain_node"`
		Contract      []hedgex      `json:"contract"`
		IpLimit       int           `json:"ip_limit"`
		test          test          `json:"test"`
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
	ChainNode = all.ChainNode
	Contract = all.Contract
	if MaxKlineCount < 1 {
		log.Panic("max kline count must > 0")
	}

	Test = all.test
}
