package service

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"hedgex-public/kline"
	"hedgex-public/model"
	"log"
	"strconv"
	"sync/atomic"
	"time"
)

var (
	ChainNodeErr     map[string]*int64
	ContractPriceErr map[string]*int64
)

func init() {
	ChainNodeErr = make(map[string]*int64)
	ContractPriceErr = make(map[string]*int64)
	for addr := range config.Contract {
		b := int64(0)
		ChainNodeErr[addr] = &b
		c := int64(0)
		ContractPriceErr[addr] = &c
	}
}

func StartRealKline() {
	gl.InitKlineData()

	// load kline data from database
	loadHistoryKline()

	// update the kline data real time from contract of blockchain network
	for addr := range config.Contract {
		go runKlineUpdate(addr)
	}
}

func runKlineUpdate(conAdd string) {
	addr, _ := strconv.ParseInt(conAdd[0:6], 0, 64)
	if addr == 0 {
		log.Panic("Get contract address error.")
	}
	b := ChainNodeErr[conAdd]
	c := ContractPriceErr[conAdd]
	watchTime := config.Contract[conAdd].WatchTime
	preTime := time.Now().Unix() - watchTime - 1
	var prePrice int64
	ticker := time.NewTicker(time.Second * config.WsTick)
	for range ticker.C {
		price, err := gl.GetIndexPrice(conAdd)
		if err != nil {
			atomic.StoreInt64(b, addr)
			gl.OutLogger.Error("%s : %v", conAdd, err)
			continue
		}
		atomic.StoreInt64(b, 0)
		if prePrice != price {
			prePrice = price
			preTime = time.Now().Unix()
		}
		if (time.Now().Unix() - preTime) > watchTime {
			atomic.StoreInt64(c, addr)
		} else {
			atomic.StoreInt64(c, 0)
		}
		updateKline(conAdd, price)
	}
}

func loadHistoryKline() {
	for conAddr := range config.Contract {
		for _, t := range kline.KlineTypes {
			if candles, err := model.GetKlineData(conAddr, t, config.MaxKlineCount); err != nil {
				log.Panic(err)
			} else {
				l := len(candles) - 1
				for j := range candles {
					kline.DefaultDrivers[conAddr].Append(t, candles[l-j])
				}
			}
		}
	}
}

//updateKline update the current kline's price
func updateKline(contract string, price int64) {
	kd := kline.DefaultDrivers[contract]
	for i := range gl.KlineTypes {
		candle, err := kd.GetCurrent(gl.KlineTypes[i])
		if err != nil {
			gl.OutLogger.Error("Get current kline error. %s : %s :%v", contract, gl.KlineTypes[i], err)
			continue
		}
		ts := time.Now().Unix() / gl.KlineTimeCount[gl.KlineTypes[i]] * gl.KlineTimeCount[gl.KlineTypes[i]]
		bChange := false
		if ts == candle[4] {
			if candle[3] != price {
				candle[3] = price
				bChange = true
			}
			if price > candle[1] {
				candle[1] = price
				bChange = true
			} else if price < candle[2] {
				candle[2] = price
				bChange = true
			}
		} else {
			bChange = true
			candle[0] = price
			candle[1] = price
			candle[2] = price
			candle[3] = price
			candle[4] = ts
		}
		if err := kd.Append(gl.KlineTypes[i], candle); err != nil {
			gl.OutLogger.Error("Append kline error. %s : %s : %v", contract, gl.KlineTypes[i], err)
		}
		if bChange {
			//store this candle to database
			if err := model.ReplaceKlineData(contract, gl.KlineTypes[i], candle); err != nil {
				gl.OutLogger.Error("replace into kline error. %v", err)
			}
		}
	}
}
