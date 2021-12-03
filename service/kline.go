package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"log"
	"time"
)

func StartRealIndexPrice() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()

	// load kline data from database
	loadHistoryKline()

	// update the kline data real time from contract of blockchain network
	ticker := time.NewTicker(time.Second * config.WsTick)
	for {
		select {
		case <-ticker.C:
			for i := range config.Contract {
				if price, err := gl.Contracts[config.Contract[i].Address].GetLatestPrice(nil); err != nil {
					gl.OutLogger.Error("Get price from contract error. ", err)
				} else {
					updateKline(config.Contract[i].Address, price.Int64())
				}
			}
		case <-QuitKline:
			ticker.Stop()
			gl.OutLogger.Info("Kline Update Service Stoped!")
			return
		}
	}
}

func loadHistoryKline() {
	for i := range config.Contract {
		klineTypes := []string{"m1", "m5", "m10", "m15", "m30", "h1", "h2", "h4", "h6", "h12", "d1"}
		for _, t := range klineTypes {
			candles, err := model.GetKlineData(config.Contract[i].Address, t, config.MaxKlineCount)
			if err != nil {
				log.Panic(err)
			}
			l := len(candles) - 1
			for j := range candles {
				gl.CurrentKlineDatas[config.Contract[i].Address].Append(t, candles[l-j])
			}
		}
	}
}

// updateKline update the current kline's price
func updateKline(contract string, price int64) {
	for i := range gl.KlineTypes {
		if _, exist := gl.CurrentKlineDatas[contract]; !exist {
			gl.CurrentKlineDatas[contract] = &gl.SafeKlineData{
				Data: make(map[string][][5]int64),
			}
		}
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
