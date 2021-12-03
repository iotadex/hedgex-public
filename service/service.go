package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/tools"
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
	if config.Service&0x1 > 0 {
		StartPublicService()
	}

	if config.Service&0x2 > 0 {
		key := tools.InputKey()
		pk := tools.AesCBCDecrypt(config.PrivateKey, key)
		gl.SetPrivateKey(pk)
		StartPrivateService()
	}

	//start listening the event of contracts
	for _, contact := range config.Contract {
		go StartFilterEvents(contact.Address)
	}
}

func StartPublicService() {
	go StartRealIndexPrice()
}

func StartPrivateService() {
	go StartExplosiveDetectServer()

	go StartExplosiveReCheck()

	go StartTakeInterestServer()
}

func Stop() {
	//stop the event listening service
	for _, contract := range config.Contract {
		QuitEvent[contract.Address] <- 1
	}

	if config.Service&0x1 > 0 {
		QuitKline <- 1
	}

	if config.Service&0x2 > 0 {
		QuitExplosiveDetect <- 1
		QuitExplosiveReCheck <- 1
	}
	ServiceWaitGroup.Wait()
}
