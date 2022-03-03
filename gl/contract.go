package gl

import (
	"context"
	"crypto/ecdsa"
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
	Contracts map[string]*hedgex.Hedgex

	erc20TransferID []byte
	chainID         *big.Int
)

func InitContract() {
	var err error
	EthHttpsClient, err = ethclient.Dial(config.ChainNode)
	if err != nil {
		log.Panic("ChainNode : ", config.ChainNode, err)
	}

	Contracts = make(map[string]*hedgex.Hedgex)
	for addr := range config.Contract {
		Contracts[addr], err = hedgex.NewHedgex(common.HexToAddress(addr), EthHttpsClient)
		if err != nil {
			log.Panic(err)
		}
	}

	chainID, err = EthHttpsClient.NetworkID(context.Background())
	if err != nil {
		log.Panic(err)
	}

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

func GetIndexPrice(add string) (int64, error) {
	price, _, _, err := Contracts[add].GetLatestPrice(nil)
	if err != nil {
		return 0, err
	}
	return price.Int64(), err
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
