package gl

import (
	"hedgex-server/config"
	"sync"
)

var KlineTypes []string
var KlineTimeCount map[string]int64

type SafeKlineData struct {
	RwMx sync.RWMutex
	Data map[string][][5]int64
}

var CurrentKlineDatas map[string]*SafeKlineData

func (skd *SafeKlineData) Get(klineType string, count int) [][5]int64 {
	skd.RwMx.RLock()
	defer skd.RwMx.RUnlock()
	l := len(skd.Data[klineType])
	if l == 0 {
		return nil
	}
	i := l - count
	if i < 0 {
		i = 0
	}
	data := make([][5]int64, l-i)
	copy(data, skd.Data[klineType][i:])
	return data
}

func (skd *SafeKlineData) GetCurrent(klineType string) [5]int64 {
	skd.RwMx.RLock()
	defer skd.RwMx.RUnlock()
	lastIndex := len(skd.Data[klineType]) - 1
	if lastIndex < 0 {
		return [5]int64{}
	}
	data := skd.Data[klineType][lastIndex]
	return data
}

func (skd *SafeKlineData) Append(klineType string, currentData [5]int64) {
	skd.RwMx.Lock()
	defer skd.RwMx.Unlock()
	count := len(skd.Data[klineType])
	if count == 0 {
		skd.Data[klineType] = make([][5]int64, 1)
		skd.Data[klineType][0] = currentData
		return
	}
	if currentData[4] == skd.Data[klineType][count-1][4] {
		skd.Data[klineType][count-1] = currentData
	} else {
		if count >= config.MaxKlineCount {
			skd.Data[klineType] = skd.Data[klineType][1:]
		}
		skd.Data[klineType] = append(skd.Data[klineType], currentData)
	}
}

func init() {
	KlineTypes = []string{"m1", "m5", "m10", "m15", "m30", "h1", "h2", "h4", "h6", "h12", "d1"}
	CurrentKlineDatas = make(map[string]*SafeKlineData)
	for _, contract := range config.Contract.Pair {
		klineSafeData := &SafeKlineData{
			Data: make(map[string][][5]int64),
		}
		CurrentKlineDatas[contract] = klineSafeData
	}
	KlineTimeCount = make(map[string]int64)
	KlineTimeCount["m1"] = 60
	KlineTimeCount["m5"] = 300
	KlineTimeCount["m10"] = 600
	KlineTimeCount["m15"] = 900
	KlineTimeCount["m30"] = 1800
	KlineTimeCount["h1"] = 3600
	KlineTimeCount["h2"] = 7200
	KlineTimeCount["h4"] = 14400
	KlineTimeCount["h6"] = 21600
	KlineTimeCount["h12"] = 43200
	KlineTimeCount["d1"] = 86400
}
