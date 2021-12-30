package main

import (
	"context"
	"hedgex-server/config"
	"hedgex-server/daemon"
	"hedgex-server/gl"
	"hedgex-server/host"
	"hedgex-server/model"
	"hedgex-server/service"
)

func main() {
	if config.Env == "product" {
		daemon.Background("./out.log", true)
	}

	//create out and err logs in logs dir
	gl.CreateLogFiles()

	//connect to mysql database
	model.ConnectToMysql()

	//init the contracts
	gl.InitContract()

	//start contract service
	service.Start()

	//start http service
	go host.StartHttpServer()

	//wait to exit single
	daemon.WaitForKill()

	//shutdown the http service
	if gl.HttpServer != nil {
		gl.HttpServer.Shutdown(context.Background())
	}

	//stop the contract service
	service.Stop()
}
