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
	//start kline service
	StartRealKline()
}
