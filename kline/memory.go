package kline

import (
	"hedgex-public/config"
	"sync"
)

type MemoryKline struct {
	tradePair string
	rwMx      sync.RWMutex
	data      map[string][][5]int64
}

func NewMemoryKline(tp string) *MemoryKline {
	mk := MemoryKline{
		tradePair: tp,
		data:      make(map[string][][5]int64),
	}
	for _, s := range KlineTypes {
		mk.data[s] = make([][5]int64, 0)
	}
	return &mk
}

func (mk *MemoryKline) Get(kt string, count int) ([][5]int64, error) {
	mk.rwMx.RLock()
	defer mk.rwMx.RUnlock()
	l := len(mk.data[kt])
	if l == 0 {
		return nil, nil
	}
	i := l - count
	if i < 0 {
		i = 0
	}
	data := make([][5]int64, l-i)
	copy(data, mk.data[kt][i:])
	return data, nil
}

func (mk *MemoryKline) GetCurrent(kt string) ([5]int64, error) {
	mk.rwMx.RLock()
	defer mk.rwMx.RUnlock()
	lastIndex := len(mk.data[kt]) - 1
	if lastIndex < 0 {
		return [5]int64{}, nil
	}
	data := mk.data[kt][lastIndex]
	return data, nil
}

func (mk *MemoryKline) Append(kt string, currentData [5]int64) error {
	mk.rwMx.Lock()
	defer mk.rwMx.Unlock()
	count := len(mk.data[kt])
	if count == 0 {
		mk.data[kt] = make([][5]int64, 1)
		mk.data[kt][0] = currentData
		return nil
	}
	if currentData[4] == mk.data[kt][count-1][4] {
		mk.data[kt][count-1] = currentData
	} else {
		if count >= config.MaxKlineCount {
			mk.data[kt] = mk.data[kt][1:]
		}
		mk.data[kt] = append(mk.data[kt], currentData)
	}
	return nil
}
