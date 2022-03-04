package service

import (
	"hedgex-public/config"
	"hedgex-public/gl"
	"hedgex-public/model"
	"log"
	"time"
)

func StartRealKline() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()

	gl.InitKlineData()

	// load kline data from database
	loadHistoryKline()

	// update the kline data real time from contract of blockchain network
	for addr := range config.Contract {
		go runKlineUpdate(addr)
	}
}

func runKlineUpdate(conAdd string) {
	ticker := time.NewTicker(time.Second * config.WsTick)
	for range ticker.C {
		price, err := gl.GetIndexPrice(conAdd)
		if err != nil {
			gl.OutLogger.Error("Get price from contract error. %s : %v", conAdd, err)
			continue
		}
		gl.OutLogger.Info("IndexPrice : %s : %d", conAdd[2:6], price)
		updateKline(conAdd, price)
	}
}

func loadHistoryKline() {
	for addr := range config.Contract {
		klineTypes := []string{"m1", "m5", "m10", "m15", "m30", "h1", "h2", "h4", "h6", "h12", "d1"}
		for _, t := range klineTypes {
			candles, err := model.GetKlineData(addr, t, config.MaxKlineCount)
			if err != nil {
				log.Panic(err)
			}
			l := len(candles) - 1
			for j := range candles {
				gl.CurrentKlineDatas[addr].Append(t, candles[l-j])
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
