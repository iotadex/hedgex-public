package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"hedgex-server/tools"
	"log"
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
	//start listening the event of contracts
	for _, contract := range config.Contract {
		go StartFilterEvents(contract.Address)
	}

	if config.Service&0x1 > 0 {
		StartPublicService()
	}

	if config.Service&0x2 > 0 {
		//get users from database
		getHistoryUsersDataFromDb()

		key := tools.InputKey()
		pk := tools.AesCBCDecrypt(config.PrivateKey, key)
		gl.SetPrivateKey(pk)
		StartPrivateService()
	}
}

func StartPublicService() {
	go StartRealIndexPrice()
}

func StartPrivateService() {
	go StartExplosiveDetectServer()

	//go StartExplosiveReCheck()

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

func getHistoryUsersDataFromDb() {
	//load user's data from database
	for _, contract := range config.Contract {
		users, _, err := model.GetUsers(contract.Address)
		if err != nil {
			log.Panic("Get users from db error. ", err)
			return
		}
		l := len(users)
		for i := 0; i < l; i++ {
			expUserList[contract.Address].Insert(&users[i])
			interestUserList[contract.Address].update(&users[i])
		}
	}
}
