package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"sync"
	"time"
)

var interestUserList map[string]*TakeInterestList //current accounts waiting for be detected to explosive
func init() {
	interestUserList = make(map[string]*TakeInterestList)
	for _, contract := range config.Contract {
		interestUserList[contract.Address] = &TakeInterestList{
			luser: make(map[string]uint),
			suser: make(map[string]uint),
		}
	}
}

//StartExplosiveDetectServer, no blocking function
func StartTakeInterestServer() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()
	timer := time.NewTicker(config.InterestTick * time.Second)
	for {
		select {
		case <-timer.C:
			ts := time.Now().Unix()
			if (ts - ts/86400*86400) > 300 {
				continue
			}
			auth, err := getAccountAuth()
			if err != nil {
				gl.OutLogger.Error("Get auth error. %v", err)
				continue
			}
			for _, contract := range config.Contract {
				//get the current price of contract
				price, err := Contracts[contract.Address].GetLatestPrice(nil)
				if err != nil {
					gl.OutLogger.Error("Get price from contract error. ", err)
					continue
				}

				node := expUserList[contract.Address].LHead.Next
				for node != nil {
					node = explosive(auth, contract.Address, node, price.Int64(), 1)
				}
				node = expUserList[contract.Address].SHead.Next
				for node != nil {
					node = explosive(auth, contract.Address, node, price.Int64(), -1)
				}
				time.Sleep(time.Second)
			}
		case <-QuitInterestDetect:
			return
		}
	}
}

func detectSlide() {
}

type TakeInterestList struct {
	luser map[string]uint // long position user
	suser map[string]uint // short position user
	mu    sync.Mutex      //user's locker
}

func (til *TakeInterestList) update(u *model.User) {
	til.mu.Lock()
	if u.Lposition > u.Sposition {
		til.luser[u.Account] = u.InterestDay
		delete(til.suser, u.Account)
	} else if u.Lposition < u.Sposition {
		til.suser[u.Account] = u.InterestDay
		delete(til.luser, u.Account)
	} else {
		delete(til.luser, u.Account)
		delete(til.suser, u.Account)
	}
	til.mu.Unlock()
}
