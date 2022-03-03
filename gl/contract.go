package gl

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"hedgex-public/config"
	"hedgex-public/hedgex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//define the client to connect to the ethereum network
	EthHttpsClient *ethclient.Client

	//contract's instance
	Contracts map[string]map[*hedgex.Hedgex]struct{}

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

	Contracts = make(map[string]map[*hedgex.Hedgex]struct{})
	for addr := range config.Contract {
		contracts := make(map[*hedgex.Hedgex]struct{})
		for i := range clients {
			con, err := hedgex.NewHedgex(common.HexToAddress(addr), clients[i])
			if err != nil {
				log.Panic(err)
			}
			contracts[con] = struct{}{}
		}
		Contracts[addr] = contracts
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

func GetIndexPrice(add string) (int64, error) {
	for con := range Contracts[add] {
		price, _, _, err := con.GetLatestPrice(nil)
		if err == nil {
			return price.Int64(), nil
		}
		OutLogger.Error("Get index pirce from contract(%s) error. %v", add, err)
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

	gasPrice, err := EthHttpsClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	nonce, err := EthHttpsClient.PendingNonceAt(context.Background(), config.Test.PublicAddress)
	if err != nil {
		return "", err
	}
	gasLimit := uint64(3000000)
	tx := types.NewTransaction(nonce, common.HexToAddress(config.Test.Token), value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), config.Test.PrivateKey)
	if err != nil {
		return "", err
	}

	err = EthHttpsClient.SendTransaction(context.Background(), signedTx)
	return tx.Hash().Hex(), err
}
