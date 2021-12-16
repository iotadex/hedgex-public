package service

import (
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"sync/atomic"
	"time"
)

func StartForceCloseService() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()
	atomic.StoreInt32(&IsRunForceClose, 1)
	for atomic.LoadInt32(&IsRunForceClose) == 1 {
		time.Sleep(time.Second * config.IndexTick)
		for i := range config.Contract {
			state, err := gl.GetPoolState(config.Contract[i].Address)
			if err != nil {
				gl.OutLogger.Error("Get pool's state from contract error. %v", err)
				continue
			}
			if state != 2 {
				continue
			}

			auth, err := gl.GetAccountAuth()
			if err != nil {
				gl.OutLogger.Error("Get auth error when explosive pool. %v", err)
				continue
			}

			//get all the users from database
			users, _, err := model.GetUsers(config.Contract[i].Address)
			if err != nil {
				gl.OutLogger.Error("Get users from db error. %v", err)
				continue
			}
			for i := range users {
				if users[i].Lposition > 0 || users[i].Sposition > 0 {
					if err = gl.ForceClose(auth, config.Contract[i].Address, users[i].Account); err != nil {
						gl.OutLogger.Error("send forceCloseAccount to contract error. %s : %v", config.Contract[i].Address, err)
					}
				}
			}
			time.Sleep(time.Second * config.IndexTick * 10)
		}
	}
}
