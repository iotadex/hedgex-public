package service

import (
	"sync"
)

var ServiceWaitGroup sync.WaitGroup
var QuitIndexPrice chan int
var QuitKline chan int

func init() {
	QuitIndexPrice = make(chan int)
	QuitKline = make(chan int)
}

func Start() {
	//start index price service
	go StartIndexPriceService()

	//start kline service
	go StartRealKline()
}

func Stop() {
	//stop kline service
	QuitKline <- 1

	//stop the indexprice service
	QuitIndexPrice <- 1

	ServiceWaitGroup.Wait()
}
