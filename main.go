package main

import (
	"hedgex-public/config"
	"hedgex-public/daemon"
	"hedgex-public/gl"
	"hedgex-public/host"
	"hedgex-public/model"
	"hedgex-public/service"
	"time"
)

func main() {
	if config.Env == "product" {
		daemon.Background("./out.log", true)
	}

	WaitForTime(config.BeginSec)

	//create out and err logs in logs dir
	gl.CreateLogFiles()

	//connect to mysql database
	model.ConnectToMysql()

	if len(config.ChainNodes) > 0 {
		//init the contracts
		gl.InitContract()

		//start contract service
		service.StartRealKline()
	}

	//start http service
	if config.HttpPort != 0 {
		host.StartHttpServer()
	}

	//wait to exit single
	daemon.WaitForKill()
}

func WaitForTime(sec int64) {
	milSec := sec * 1000
	now := time.Now().UnixMilli() % 60000
	if now < milSec {
		time.Sleep(time.Duration(now-milSec) * time.Millisecond)
	} else {
		time.Sleep(time.Millisecond * time.Duration(60000-now+milSec))
	}
}
