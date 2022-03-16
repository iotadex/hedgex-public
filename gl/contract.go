package gl

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"hedgex-public/config"
	"hedgex-public/hedgex"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//contract's instance
	Contracts map[string]map[*hedgex.Hedgex]string

	erc20TransferID []byte
	chainID         *big.Int
)

func InitContract() {
	clients := make([]*ethclient.Client, 0, len(config.ChainNodes))
	for i := range config.ChainNodes {
		client, err := ethclient.Dial(config.ChainNodes[i])
		if err != nil {
			log.Panic("ChainNode : ", config.ChainNodes[i], err)
		}
		clients = append(clients, client)
	}

	Contracts = make(map[string]map[*hedgex.Hedgex]string)
	for conAddr := range config.Contract {
		contracts := make(map[*hedgex.Hedgex]string)
		for i := range clients {
			con, err := hedgex.NewHedgex(common.HexToAddress(conAddr), clients[i])
			if err != nil {
				log.Panic(err)
			}
			contracts[con] = config.ChainNodes[i]
		}
		Contracts[conAddr] = contracts
	}

	var err error
	chainID, err = clients[0].NetworkID(context.Background())
	if err != nil {
		log.Panic(err)
	}

	if len(config.Test.Wallet) > 0 {
		config.Test.PrivateKey, err = crypto.HexToECDSA(config.Test.Wallet)
		if err != nil {
			log.Panic("Get privatekey error.", err)
		}
		publicKey := config.Test.PrivateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Panic("error casting public key to ECDSA")
		}
		config.Test.PublicAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	}
}

func GetIndexPrice(conAddr string) (int64, error) {
	prices := make(map[int64]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for con := range Contracts[conAddr] {
		wg.Add(1)
		go func(con *hedgex.Hedgex) {
			ch := make(chan int64)
			go func() {
				if price, _, _, err := con.GetLatestPrice(nil); err != nil {
					ch <- 0
					OutLogger.Error("Get index pirce error. %s : %v", Contracts[conAddr][con], err)
				} else {
					ch <- price.Int64()
				}
			}()
			select {
			case price := <-ch:
				mu.Lock()
				prices[price]++
				mu.Unlock()
			case <-time.After(3 * time.Second):
				OutLogger.Error("Get index price over time. %s", Contracts[conAddr][con])
			}
			wg.Done()
		}(con)
	}
	wg.Wait()
	price := int64(0)
	maxCount := 0
	for p, c := range prices {
		if (c > maxCount) && (p > 0) {
			price = p
			maxCount = c
		}
	}
	if maxCount > 0 {
		OutLogger.Info("Index Price : %s : %d : %d", conAddr[2:6], maxCount, price)
		return price, nil
	}
	return 0, errors.New("get index price error")
}

func SendTestCoins(to string) (string, error) {
	paddedAddress := common.LeftPadBytes(common.HexToAddress(to).Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(config.Test.SendAmount).Bytes(), 32)

	var data []byte
	data = append(data, erc20TransferID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	value := big.NewInt(0)

	client, err := ethclient.Dial(config.ChainNodes[0])
	if err != nil {
		return "", err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	nonce, err := client.PendingNonceAt(context.Background(), config.Test.PublicAddress)
	if err != nil {
		return "", err
	}
	gasLimit := uint64(3000000)
	tx := types.NewTransaction(nonce, common.HexToAddress(config.Test.Token), value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), config.Test.PrivateKey)
	if err != nil {
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	return tx.Hash().Hex(), err
}
