package gl

import (
	"context"
	"crypto/ecdsa"
	"hedgex-public/config"
	"hedgex-public/indexprice"
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
	indexPriceContracts map[*indexprice.Indexprice]string
)

func InitContract() {
	indexPriceContracts = make(map[*indexprice.Indexprice]string)
	for i := range config.ChainNodes {
		client, err := ethclient.Dial(config.ChainNodes[i])
		if err != nil {
			log.Panic("ChainNode : ", config.ChainNodes[i], err)
		}

		con, err := indexprice.NewIndexprice(common.HexToAddress(config.IndexPriceConAddr), client)
		if err != nil {
			log.Panic(err)
		}

		indexPriceContracts[con] = config.ChainNodes[i]
	}

	if len(config.Test.Wallet) > 0 {
		var err error
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

func GetIndexPrices(conAddresses []common.Address) []int64 {
	begin := time.Now().UnixMilli()
	var wg sync.WaitGroup
	prices := make([][]*big.Int, len(indexPriceContracts))
	i := 0
	for con := range indexPriceContracts {
		wg.Add(1)
		go func(con *indexprice.Indexprice, i int) {
			ch := make(chan []*big.Int)
			go func() {
				if price, err := con.IndexPrice(nil, conAddresses); err != nil {
					ch <- nil
					OutLogger.Error("Get index pirce error. %s : %v", indexPriceContracts[con], err)
				} else {
					ch <- price
				}
			}()
			select {
			case p := <-ch:
				prices[i] = p
			case <-time.After(3 * time.Second):
				OutLogger.Error("Get index price over time. %s", indexPriceContracts[con])
			}
			wg.Done()
		}(con, i)
		i++
	}
	wg.Wait()

	count := len(conAddresses)
	statPrices := make([]map[int64]int, count)
	for i := 0; i < count; i++ {
		statPrices[i] = make(map[int64]int)
	}
	for _, ps := range prices {
		for i := 0; i < len(ps); i++ {
			statPrices[i][ps[i].Int64()]++
		}
	}

	ips := make([]int64, count)
	nodes := make([]int, count)
	for i := 0; i < count; i++ {
		price := int64(0)
		maxCount := 0
		for p, c := range statPrices[i] {
			if (c > maxCount) && (p > 0) {
				price = p
				maxCount = c
			}
		}
		ips[i] = price
		nodes[i] = maxCount
	}
	OutLogger.Info("Index prices : %v : %v : %d", ips, nodes, time.Now().UnixMilli()-begin)
	return ips
}

func SendTestCoins(to string) (string, error) {
	paddedAddress := common.LeftPadBytes(common.HexToAddress(to).Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(config.Test.SendAmount).Bytes(), 32)
	methodid := []byte{0xa9, 0x05, 0x9c, 0xbb}
	var data []byte
	data = append(data, methodid...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	value := big.NewInt(0)

	client, err := ethclient.Dial(config.ChainNodes[0])
	if err != nil {
		return "", err
	}

	chainID, err := client.NetworkID(context.Background())
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
