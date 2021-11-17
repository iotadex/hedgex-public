package service

import (
	"context"
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//StartFilterEvents start to filter the contract's events
func StartFilterEvents(contractAddress string) {
	gl.ServiceWaitGroup.Add(1)
	GetHistoryEventLogs(contractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}
StartFilter:
	logs := make(chan types.Log)
	sub, err := gl.EthWssClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Panic("Get event logs from eth wss client error. ", err)
	}

	for {
		select {
		case <-gl.QuitChan[contractAddress]:
			goto ForEnd
		case err := <-sub.Err():
			gl.OutLogger.Error("Evnet wss sub error. %v", err)
			gl.OutLogger.Warn("The EthWssClient will be redialed...")
			time.Sleep(time.Second * 10)
			gl.EthWssClient, err = ethclient.Dial(config.Contract.Wss)
			if err != nil {
				gl.OutLogger.Warn("The EthWssClient redial error. %v", err)
				continue
			}
			goto StartFilter
		case vLog := <-logs:
			dealEventLog(contractAddress, &vLog)
		}
	}
ForEnd:
	gl.EthWssClient.Close()
	gl.OutLogger.Info("Events Filter(%s) Service Stoped!", contractAddress)
	gl.ServiceWaitGroup.Done()
}

func GetHistoryEventLogs(contractAddress string) {
	//Get the last block number from database
	lastBlock, err := model.GetLastBlock(contractAddress)
	if err != nil {
		log.Panic("Get the last block from database error. ", err)
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(lastBlock),
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}

	logs, err := gl.EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Panic("Get event logs from eth client error. ", err)
	}

	for _, log := range logs {
		dealEventLog(contractAddress, &log)
	}
}

func dealEventLog(contract string, vLog *types.Log) {
	if len(vLog.Topics) > 0 {
		tx := vLog.TxHash.Hex()
		block := vLog.BlockNumber
		data, err := gl.ContractAbi.Unpack(gl.EventNames[vLog.Topics[0].Hex()], vLog.Data)
		if err != nil {
			gl.OutLogger.Error("Unpack the event log error. tx(%s) : %s : %s", tx, gl.EventNames[vLog.Topics[0].Hex()], err.Error())
			return
		}
		account := common.HexToAddress(vLog.Topics[1].Hex()).Hex()
		switch vLog.Topics[0].Hex() {
		case gl.MintEvent:
			amount := data[0].(*big.Int).Uint64()
			err = model.InsertMint(tx, contract, account, amount, block)
		case gl.BurnEvent:
			amount := data[0].(*big.Int).Uint64()
			err = model.InsertBurn(tx, contract, account, amount, block)
		case gl.RechargeEvent:
			amount := data[0].(*big.Int).Uint64()
			err = model.InsertRecharge(tx, contract, account, amount, block)
			updateUser(contract, account, block)
		case gl.WithdrawEvent:
			amount := data[0].(*big.Int).Uint64()
			err = model.InsertWithdraw(tx, contract, account, amount, block)
			updateUser(contract, account, block)
		case gl.TradeEvent:
			direction := data[0].(int8)
			amount := data[1].(*big.Int).Uint64()
			price := data[2].(*big.Int).Uint64()
			err = model.InsertTrade(tx, contract, account, direction, amount, price, block)
			updateUser(contract, account, block)
		case gl.ExplosiveEvent:
			direction := data[0].(int8)
			amount := data[1].(*big.Int).Uint64()
			price := data[2].(*big.Int).Uint64()
			err = model.InsertTrade(tx, contract, account, direction, amount, price, block)
			updateUser(contract, account, block)
		case gl.TakeInterestEvent:
			direction := data[0].(int8)
			amount := data[1].(*big.Int).Uint64()
			price := data[2].(*big.Int).Uint64()
			err = model.InsertInterest(tx, contract, account, direction, amount, price, block)
			updateUser(contract, account, block)
		}
		if err != nil {
			gl.OutLogger.Error("insert into database error. %s : %s", gl.EventNames[vLog.Topics[0].Hex()], err.Error())
		}
	}
}

func updateUser(contract string, account string, block uint64) {
	trader, err := gl.Contracts[contract].Traders(nil, common.HexToAddress(account))
	if err != nil {
		gl.OutLogger.Error("Get account's position data from blockchain error. %s", err.Error())
	}
	user := gl.User{
		Margin:    trader.Margin.Int64(),
		Lposition: trader.LongAmount.Uint64(),
		Lprice:    trader.LongPrice.Uint64(),
		Sposition: trader.ShortAmount.Uint64(),
		Sprice:    trader.ShortPrice.Uint64(),
		Block:     block,
	}
	if err := model.UpdateUser(contract, account, user); err != nil {
		gl.OutLogger.Error("Update account's data in database error. %s", err.Error())
	}
}
