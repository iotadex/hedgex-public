package service

import (
	"hedgex-server/config"
	"sync"
)

var ServiceWaitGroup sync.WaitGroup
var QuitKline chan int
var QuitEvent map[string]chan int
var QuitExplosiveDetect chan int
var QuitExplosiveReCheck chan int
var QuitInterestDetect chan int

func init() {
	QuitKline = make(chan int)
	QuitEvent = make(map[string]chan int)
	QuitExplosiveDetect = make(chan int)
	QuitExplosiveReCheck = make(chan int)
}

func Start() {
	if config.Service == 0 {
		//start geting index price for general kline data by real
		go StartRealIndexPrice()
	} else {
		//start detecting the explosive account
		go StartExplosiveDetectServer()

		go StartExplosiveReCheck()
	}

	//start listening the event of contracts
	for _, contact := range config.Contract {
		go StartFilterEvents(contact.Address)
	}
}

func Stop() {
	for _, contract := range config.Contract {
		QuitEvent[contract.Address] <- 1
	}
	if config.Service == 0 {
		QuitKline <- 1
	} else {
		QuitExplosiveDetect <- 1
		QuitExplosiveReCheck <- 1
	}
	ServiceWaitGroup.Wait()
}
