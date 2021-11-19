package service

import (
	"context"
	"crypto/ecdsa"
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/host"
	"hedgex-server/tools"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

var ServiceWaitGroup sync.WaitGroup
var QuitKline chan int
var QuitEvent map[string]chan int
var QuitExplosiveDetect chan int
var QuitExplosiveReCheck chan int

func init() {
	QuitKline = make(chan int)
	QuitEvent = make(map[string]chan int)
	QuitExplosiveDetect = make(chan int)
	QuitExplosiveReCheck = make(chan int)
}

func Start() {
	if config.Service == 0 {
		createPrivateKey()

		//start host server
		go host.StartHttpServer()

		//start geting index price for general kline data by real
		go StartRealIndexPrice()
	} else {
		//start detecting the explosive account
		go StartExplosiveDetectServer()

		go StartExplosiveReCheck()
	}

	//start listening the event of contracts
	for _, add := range config.Contract.Pair {
		go StartFilterEvents(add)
	}
}

func Stop() {
	for _, add := range config.Contract.Pair {
		QuitEvent[add] <- 1
	}
	if config.Service == 0 {
		QuitKline <- 1
	} else {
		QuitExplosiveDetect <- 1
		QuitExplosiveReCheck <- 1
	}

	if gl.HttpServer != nil {
		gl.HttpServer.Shutdown(context.Background())
	}
	ServiceWaitGroup.Wait()
}

func createPrivateKey() {
	var err error
	key := tools.InputKey()
	privateKey, err = crypto.HexToECDSA(tools.AesCBCDecrypt(config.PrivateKey, key))
	if err != nil {
		log.Panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
}
