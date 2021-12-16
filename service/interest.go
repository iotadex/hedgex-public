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
			luser: make(map[string]*interestUser),
			suser: make(map[string]*interestUser),
		}
	}
}

//StartExplosiveDetectServer, no blocking function
func StartTakeInterestServer() {
	for {
		ts := time.Now().Unix()
		dayCount := ts / 86400
		offset := ts - dayCount*86400
		if offset <= config.Interest.Begin || offset > config.Interest.End {
			sleepTime := (86400 - offset) / 2
			if offset <= config.Interest.Begin {
				sleepTime = config.Interest.Begin - offset + 1
			} else if sleepTime < int64(config.Interest.Tick) {
				sleepTime = int64(config.Interest.Tick)
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
			continue
		}
		auth, err := gl.GetAccountAuth()
		if err != nil {
			gl.OutLogger.Error("Get auth error. %v", err)
			continue
		}

		ServiceWaitGroup.Add(1)
		for _, contract := range config.Contract {
			//get the pool's position
			_, lp, _, sp, _, _, err := gl.GetPoolPosition(contract.Address)
			if err != nil {
				gl.OutLogger.Error("Get account's position data from blockchain error. %s : %v", contract.Address, err)
				continue
			}
			til := interestUserList[contract.Address]
			var l map[string]*interestUser
			if lp > sp {
				l = til.getShortUsers()
			} else if lp < sp {
				l = til.getLongUsers()
			}

			for k, v := range l {
				if v.day < uint(dayCount) {
					if gl.DetectSlide(auth, contract.Address, k) == nil {
						gl.OutLogger.Info("send interest over. %s", k)
						til.flag(k, uint(dayCount))
					}
				}
			}
		}
		ServiceWaitGroup.Done()
	}
}

type interestUser struct {
	block uint64
	day   uint
}

type TakeInterestList struct {
	luser map[string]*interestUser // long position user
	suser map[string]*interestUser // short position user
	mu    sync.Mutex               //user's locker
}

func (til *TakeInterestList) flag(account string, day uint) {
	til.mu.Lock()
	defer til.mu.Unlock()
	if _, exist := til.luser[account]; exist {
		til.luser[account].day = day
	} else if _, exist := til.suser[account]; exist {
		til.suser[account].day = day
	}
}

func (til *TakeInterestList) update(u *model.User) {
	til.mu.Lock()
	defer til.mu.Unlock()
	v, exist := til.luser[u.Account]
	if !exist {
		v, exist = til.suser[u.Account]
	}
	if exist && (v.block > u.Block) {
		return
	}
	if u.Lposition > u.Sposition {
		til.luser[u.Account] = &interestUser{u.Block, u.InterestDay}
		delete(til.suser, u.Account)
	} else if u.Lposition < u.Sposition {
		til.suser[u.Account] = &interestUser{u.Block, u.InterestDay}
		delete(til.luser, u.Account)
	} else {
		delete(til.luser, u.Account)
		delete(til.suser, u.Account)
	}
}

func (til *TakeInterestList) getShortUsers() map[string]*interestUser {
	su := make(map[string]*interestUser)
	til.mu.Lock()
	defer til.mu.Unlock()
	m := til.suser
	for k, v := range m {
		su[k] = v
	}
	return su
}

func (til *TakeInterestList) getLongUsers() map[string]*interestUser {
	lu := make(map[string]*interestUser)
	til.mu.Lock()
	defer til.mu.Unlock()
	m := til.luser
	for k, v := range m {
		lu[k] = v
	}
	return lu
}
