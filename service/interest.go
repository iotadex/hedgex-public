package service

import (
	"context"
	"hedgex-server/config"
	"hedgex-server/contract/hedgex"
	"hedgex-server/gl"
	"hedgex-server/model"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
			_lp, _, _sp, _, err := gl.Contracts[contract.Address].GetPoolPosition(nil)
			if err != nil {
				gl.OutLogger.Error("Get account's position data from blockchain error. %s", err.Error())
				continue
			}
			til := interestUserList[contract.Address]
			var d int8 = 1
			if _lp.Uint64() < _sp.Uint64() {
				d = -1
			}
			l := til.getList(d)
			for k, v := range l {
				if v < uint(dayCount) {
					if detectSlide(auth, gl.Contracts[contract.Address], k) {
						til.flag(d, k, uint(dayCount))
					}
				}
			}
		}
		ServiceWaitGroup.Done()
	}
}

func detectSlide(auth *bind.TransactOpts, contract *hedgex.Hedgex, account string) bool {
	nonce, err := gl.EthHttpsClient.PendingNonceAt(context.Background(), gl.PublicAddress)
	if err != nil {
		gl.OutLogger.Error("Take interest : Get nonce error address(%s). %v", gl.PublicAddress, err)
		return false
	}
	auth.Nonce = big.NewInt(int64(nonce))
	if _, err := contract.DetectSlide(auth, common.HexToAddress(account), common.HexToAddress(config.Interest.ToAddress)); err != nil {
		gl.OutLogger.Error("Transaction with detect slide error. %v", err)
		return false
	}
	return true
}

type TakeInterestList struct {
	luser map[string]uint // long position user
	suser map[string]uint // short position user
	mu    sync.Mutex      //user's locker
}

func (til *TakeInterestList) flag(d int8, account string, day uint) {
	til.mu.Lock()
	m := til.luser
	if d < 0 {
		m = til.suser
	}
	if _, exist := m[account]; exist {
		m[account] = day
	}
	til.mu.Unlock()
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

func (til *TakeInterestList) getList(d int8) map[string]uint {
	l := make(map[string]uint)
	til.mu.Lock()
	m := til.luser
	if d < 0 {
		m = til.suser
	}
	for k, v := range m {
		l[k] = v
	}
	til.mu.Unlock()
	return l
}
