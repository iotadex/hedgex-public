package gl

import (
	"context"
	"crypto/ecdsa"
	"hedgex-public/config"
	"hedgex-public/contract/hedgex"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//define the client to connect to the ethereum network
	EthHttpsClient *ethclient.Client

	//define the contract's abi
	ContractAbi abi.ABI

	//definde the hash string of contract's event
	MintEvent         string
	BurnEvent         string
	RechargeEvent     string
	WithdrawEvent     string
	TradeEvent        string
	ExplosiveEvent    string
	TakeInterestEvent string
	TransferEvent     string
	EventNames        map[string]string

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

	ContractAbi, err = abi.JSON(strings.NewReader(string(hedgex.HedgexABI)))
	if err != nil {
		log.Panic(err)
	}

	MintEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Mint"].Sig)).Hex()
	BurnEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Burn"].Sig)).Hex()
	RechargeEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Recharge"].Sig)).Hex()
	WithdrawEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Withdraw"].Sig)).Hex()
	TradeEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Trade"].Sig)).Hex()
	ExplosiveEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Explosive"].Sig)).Hex()
	TakeInterestEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["TakeInterest"].Sig)).Hex()
	TransferEvent = crypto.Keccak256Hash([]byte(ContractAbi.Events["Transfer"].Sig)).Hex()

	EventNames = make(map[string]string)
	EventNames[MintEvent] = "Mint"
	EventNames[BurnEvent] = "Burn"
	EventNames[RechargeEvent] = "Recharge"
	EventNames[WithdrawEvent] = "Withdraw"
	EventNames[TradeEvent] = "Trade"
	EventNames[ExplosiveEvent] = "Explosive"
	EventNames[TakeInterestEvent] = "TakeInterest"
	EventNames[TransferEvent] = "Transfer"

	erc20TransferID = []byte{0xa9, 0x05, 0x9c, 0xbb} //transfer(address,uint256)

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

func GetPoolPosition(add string) (int64, int64, int64, int64, int64, uint8, error) {
	_total, _lp, _lprice, _sp, _sprice, _state, err := Contracts[add].GetPoolPosition(nil)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}
	return _total.Int64(), _lp.Int64(), _lprice.Int64(), _sp.Int64(), _sprice.Int64(), _state, nil
}

func GetPoolState(add string) (uint8, error) {
	return Contracts[add].PoolState(nil)
}

func GetCurrentBlockNumber() (uint64, error) {
	return EthHttpsClient.BlockNumber(context.Background())
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
