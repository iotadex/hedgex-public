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

type redis struct {
	Addr string `json:"addr"`
}

type TradePair struct {
	MarginCoin string `json:"margin_coin"`
	TradeCoin  string `json:"trade_coin"`
	WatchTime  int64  `json:"watch_time"`
	Params     Param  `json:"param"`
}

type Param struct {
	Leverage                    int     `json:"leverage"`
	MinAmount                   float64 `json:"min_amount"`
	KeepMarginScale             int     `json:"keep_margin_scale"`
	FeeRate                     float64 `json:"fee_rate"`
	SingleCloseLimitRate        float64 `json:"single_close_limit_rate"`
	SingleOpenLimitRate         float64 `json:"single_open_limit_rate"`
	PoolNetAmountRateLimitOpen  float64 `json:"r_open"`
	PoolNetAmountRateLimitPrice float64 `json:"r_price"`
	Token0                      string  `json:"token0"`
	Token0Decimal               int     `json:"token0_decimal"`
	DailyInterestRateBase       float64 `json:"daily_interest_rate_base"`
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
	Env               string
	Db                db
	Redis             redis
	HttpPort          int
	BeginSec          int64
	WsTick            time.Duration
	MaxKlineCount     int
	IndexPriceConAddr string
	ChainNodes        []string
	Contract          map[string]TradePair
	IpLimit           int
	Test              test
)

//Load load config file
func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	type Config struct {
		Env               string               `json:"env"`
		HttpPort          int                  `json:"http_port"`
		BeginSec          int64                `json:"begin_sec"`
		WsTick            time.Duration        `json:"ws_tick"`
		KlineMaxCount     int                  `json:"kline_max_count"`
		Db                db                   `json:"db"`
		Redis             redis                `json:"redis"`
		IndexPriceConAddr string               `json:"index_price"`
		ChainNodes        []string             `json:"chain_node"`
		Contract          map[string]TradePair `json:"contract"`
		IpLimit           int                  `json:"ip_limit"`
		Test              test                 `json:"test"`
	}
	all := &Config{}
	if err = json.NewDecoder(file).Decode(all); err != nil {
		log.Panic(err)
	}
	Env = all.Env
	Db = all.Db
	Redis = all.Redis
	HttpPort = all.HttpPort
	WsTick = all.WsTick
	MaxKlineCount = all.KlineMaxCount
	IndexPriceConAddr = all.IndexPriceConAddr
	ChainNodes = all.ChainNodes
	Contract = all.Contract
	if MaxKlineCount < 1 {
		log.Panic("max kline count must > 0")
	}
	for _, tp := range Contract {
		if tp.WatchTime < 1 {
			log.Panic("Watch time must > 0")
		}
	}

	Test = all.Test
}
