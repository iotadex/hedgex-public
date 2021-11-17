package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"log"
	"sync/atomic"
	"time"
)

func StartRealIndexPrice() {
	gl.ServiceWaitGroup.Add(1)
	// load kline data from database
	for i := range config.Contract.Pair {
		klineTypes := []string{"m1", "m5", "m10", "m15", "m30", "h1", "h2", "h4", "h6", "h12", "d1"}
		for _, t := range klineTypes {
			candles, err := model.GetKlineData(config.Contract.Pair[i], t, config.MaxKlineCount)
			if err != nil {
				log.Panic(err)
			}
			l := len(candles) - 1
			for j := range candles {
				gl.CurrentKlineDatas[config.Contract.Pair[i]].Append(t, candles[l-j])
			}
		}
	}

	// update the kline data real time from contract of blockchain network56918951
	if atomic.LoadInt32(&gl.KLineServerIsRun) == 1 {
		return
	}
	atomic.StoreInt32(&gl.KLineServerIsRun, 1)
	ticker := time.NewTicker(time.Second * config.WsTick)
	defer ticker.Stop()
	for range ticker.C {
		if atomic.LoadInt32(&gl.KLineServerIsRun) == 0 {
			ticker.Stop()
			break
		}
		for i := range config.Contract.Pair {
			if price, err := gl.Contracts[config.Contract.Pair[i]].GetLatestPrice(nil); err != nil {
				gl.OutLogger.Error("Get price from contract error. ", err)
			} else {
				updateKline(config.Contract.Pair[i], price.Int64())
			}
		}
	}
	gl.OutLogger.Info("Kline Update Service Stoped!")
	gl.ServiceWaitGroup.Done()
}

// updateKline update the current kline's price
func updateKline(contract string, price int64) {
	for i := range gl.KlineTypes {
		candle := gl.CurrentKlineDatas[contract].GetCurrent(gl.KlineTypes[i])
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
		gl.CurrentKlineDatas[contract].Append(gl.KlineTypes[i], candle)
		if bChange {
			//store this candle to database
			if err := model.ReplaceKlineData(contract, gl.KlineTypes[i], candle); err != nil {
				gl.OutLogger.Error("replace into kline error. %v", err)
			}
		}
	}
}

// GetLatestPrice unused func
func GetLatestPrice(address string) int64 {
	return 0
}
