package service

import (
	"crypto/ecdsa"
	"hedgex-server/config"
	"hedgex-server/contract/hedgex"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	//define the client to connect to the ethereum network
	EthHttpsClient *ethclient.Client
	EthWssClient   *ethclient.Client

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
	EventNames        map[string]string

	//contract's instance
	Contracts map[string]*hedgex.Hedgex

	//
	privateKey    *ecdsa.PrivateKey
	publicAddress common.Address
)

func init() {
	var err error

	EthHttpsClient, err = ethclient.Dial(config.Contract.Https)
	if err != nil {
		log.Panic(err)
	}

	Contracts = make(map[string]*hedgex.Hedgex)
	for i := range config.Contract.Pair {
		Contracts[config.Contract.Pair[i]], err = hedgex.NewHedgex(common.HexToAddress(config.Contract.Pair[i]), EthHttpsClient)
		if err != nil {
			log.Panic(err)
		}
	}

	EthWssClient, err = ethclient.Dial(config.Contract.Wss)
	if err != nil {
		log.Panic(err)
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

	EventNames = make(map[string]string)
	EventNames[MintEvent] = "Mint"
	EventNames[BurnEvent] = "Burn"
	EventNames[RechargeEvent] = "Recharge"
	EventNames[WithdrawEvent] = "Withdraw"
	EventNames[TradeEvent] = "Trade"
	EventNames[ExplosiveEvent] = "Explosive"
	EventNames[TakeInterestEvent] = "TakeInterest"
}
