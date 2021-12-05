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
	//load user's data from database
	for _, contract := range config.Contract {
		users, _, err := model.GetUsers(contract.Address)
		if err != nil {
			gl.OutLogger.Error("Get users from db error. %v", err)
			return
		}
		l := len(users)
		for i := 0; i < l; i++ {
			interestUserList[contract.Address].update(&users[i])
		}
	}

	for {
		ts := time.Now().Unix()
		dayCount := ts / 86400
		offset := ts - dayCount*86400
		if offset <= config.Interest.Begin || offset > config.Interest.End {
			left := 86400 - offset
			sleepTime := left / 2
			if sleepTime < int64(config.Interest.Tick) {
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
			lp, sp, err := gl.GetPoolPosition(contract.Address)
			if err != nil {
				continue
			}
			til := interestUserList[contract.Address]
			var d int8 = 1
			if lp < sp {
				d = -1
			}
			l := til.getList(d)
			for k, v := range l {
				if v.day < uint(dayCount) {
					if gl.DetectSlide(auth, contract.Address, k) == nil {
						gl.OutLogger.Info("send interest over. %s", k)
						til.flag(d, k, uint(dayCount))
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

func (til *TakeInterestList) flag(d int8, account string, day uint) {
	til.mu.Lock()
	defer til.mu.Unlock()
	m := til.luser
	if d < 0 {
		m = til.suser
	}
	if _, exist := m[account]; exist {
		m[account].day = day
	}
}

func (til *TakeInterestList) update(u *model.User) {
	til.mu.Lock()
	defer til.mu.Unlock()
	v, exist := til.luser[u.Account]
	if !exist {
		v, exist = til.suser[u.Account]
	}
	if exist {
		if v.block > u.Block {
			return
		}
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

func (til *TakeInterestList) getList(d int8) map[string]*interestUser {
	l := make(map[string]*interestUser)
	til.mu.Lock()
	defer til.mu.Unlock()
	m := til.luser
	if d < 0 {
		m = til.suser
	}
	for k, v := range m {
		l[k] = v
	}
	return l
}
