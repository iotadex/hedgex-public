package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"sync/atomic"
	"time"
)

var IndexPrices map[string]*int64

func init() {
	IndexPrices = make(map[string]*int64)
	for i := range config.Contract {
		IndexPrices[config.Contract[i].Address] = new(int64)
	}
}

func StartIndexPriceService() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()
	ticker := time.NewTicker(time.Second * config.IndexTick)
	for {
		select {
		case <-ticker.C:
			for i := range config.Contract {
				if price, err := gl.GetIndexPrice(config.Contract[i].Address); err != nil {
					gl.OutLogger.Error("Get price from contract error. ", err)
				} else {
					atomic.StoreInt64(IndexPrices[config.Contract[i].Address], price)
				}
			}
		case <-QuitIndexPrice:
			ticker.Stop()
			gl.OutLogger.Info("Index Price Update Service Stoped!")
			return
		}
	}
}
