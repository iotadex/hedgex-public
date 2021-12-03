package main

import (
	"context"
	"hedgex-server/config"
	"hedgex-server/daemon"
	"hedgex-server/gl"
	"hedgex-server/host"
	"hedgex-server/service"
)

func main() {
	daemon.Background("./out.log", true)

	service.Start()
	if config.Service == 0 {
		//start host server
		go host.StartHttpServer()
	}

	daemon.WaitForKill()

	if gl.HttpServer != nil {
		gl.HttpServer.Shutdown(context.Background())
	}
	service.Stop()
}
