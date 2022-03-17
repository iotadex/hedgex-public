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

	"github.com/ethereum/go-ethereum/common"
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
	// load kline data from database
	loadHistoryKline()

	// update the kline data real time from contract of blockchain network
	preTimes := make([]int64, 0)
	watchTimes := make([]int64, 0)
	addressesHex := make([]string, 0)
	addresses := make([]common.Address, 0)
	addressesInt64 := make([]int64, 0)
	for conAddr := range config.Contract {
		addresses = append(addresses, common.HexToAddress(conAddr))
		addressesHex = append(addressesHex, conAddr)
		addr, _ := strconv.ParseInt(conAddr[0:10], 0, 64)
		addressesInt64 = append(addressesInt64, addr)
		watchTimes = append(watchTimes, config.Contract[conAddr].WatchTime)
		preTimes = append(preTimes, time.Now().Unix()-config.Contract[conAddr].WatchTime-1)
	}
	count := len(addresses)

	prePrices := make([]int64, count)
	ticker := time.NewTicker(time.Second * config.WsTick)
	for range ticker.C {
		prices := gl.GetIndexPrices(addresses)
		for i := 0; i < count; i++ {
			if prices[i] != 0 {
				atomic.StoreInt64(ContractPriceErr[addressesHex[i]], 0)
				updateKline(addressesHex[i], prices[i])
			} else {
				atomic.StoreInt64(ChainNodeErr[addressesHex[i]], addressesInt64[i])
				continue
			}
			if prePrices[i] != prices[i] {
				prePrices[i] = prices[i]
				preTimes[i] = time.Now().Unix()
				atomic.StoreInt64(ContractPriceErr[addressesHex[i]], 0)
			} else if (time.Now().Unix() - preTimes[i]) > watchTimes[i] {
				atomic.StoreInt64(ContractPriceErr[addressesHex[i]], addressesInt64[i])
			}
		}
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
	for i := range kline.KlineTypes {
		candle, err := kd.GetCurrent(kline.KlineTypes[i])
		if err != nil {
			gl.OutLogger.Error("Get current kline error. %s : %s :%v", contract, kline.KlineTypes[i], err)
			continue
		}
		ts := time.Now().Unix() / kline.KlineTimeCount[kline.KlineTypes[i]] * kline.KlineTimeCount[kline.KlineTypes[i]]
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
		if err := kd.Append(kline.KlineTypes[i], candle); err != nil {
			gl.OutLogger.Error("Append kline error. %s : %s : %v", contract, kline.KlineTypes[i], err)
		}
		if bChange {
			//store this candle to database
			if err := model.ReplaceKlineData(contract, kline.KlineTypes[i], candle); err != nil {
				gl.OutLogger.Error("replace into kline error. %v", err)
			}
		}
	}
}
