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
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()
	QuitEvent[contractAddress] = make(chan int)

	getHistoryEventLogs(contractAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}
StartFilter:
	eventLogChan := make(chan types.Log)
	sub, err := gl.EthWssClient.SubscribeFilterLogs(context.Background(), query, eventLogChan)
	if err != nil {
		log.Panic("Get event logs from eth wss client error. ", contractAddress, err)
	}
	for {
		select {
		case <-QuitEvent[contractAddress]:
			gl.EthWssClient.Close()
			gl.OutLogger.Info("Events Filter(%s) Service Stoped!", contractAddress)
			return
		case err := <-sub.Err():
			gl.OutLogger.Error("Evnet wss sub error. %v", err)
			gl.OutLogger.Warn("The EthWssClient will be redialed...")
			time.Sleep(time.Second * 10)
			gl.EthWssClient, err = ethclient.Dial(config.ChainNode.Wss)
			if err != nil {
				gl.OutLogger.Warn("The EthWssClient redial error. %v", err)
				continue
			}
			goto StartFilter
		case vLog := <-eventLogChan:
			dealEventLog(contractAddress, &vLog)
		}
	}
}

func getHistoryEventLogs(contractAddress string) {
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
		var err error
		var data []interface{}
		topic0 := vLog.Topics[0].Hex()
		event, exist := gl.EventNames[topic0]
		if exist {
			if len(vLog.Data) > 0 {
				data, err = gl.ContractAbi.Unpack(event, vLog.Data)
				if err != nil {
					gl.OutLogger.Error("Unpack the event log error. tx(%s) : %s : %s : %s", tx, event, gl.EventNames[topic0], err.Error())
					return
				}
			}
		}
		account := common.HexToAddress(vLog.Topics[1].Hex()).Hex()
		switch topic0 {
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
			err = model.InsertExplosive(tx, contract, account, direction, amount, price, block)
			updateUser(contract, account, block)
		case gl.TakeInterestEvent:
			direction := data[0].(int8)
			amount := data[1].(*big.Int).Uint64()
			price := data[2].(*big.Int).Uint64()
			err = model.InsertInterest(tx, contract, account, direction, amount, price, block)
			updateUser(contract, account, block)
		case gl.TransferEvent:
			gl.OutLogger.Info("Transfer from %s to %s : %v", account, vLog.Topics[2].Hex(), data[0])
		default:
			gl.OutLogger.Info("Topics : %v   ;   data  :  %v", vLog.Topics, vLog.Data)
		}
		if err != nil {
			gl.OutLogger.Error("insert into database error. %s : %s", gl.EventNames[vLog.Topics[0].Hex()], err.Error())
		}
	}
}

func updateUser(contract string, account string, block uint64) {
	trader, err := gl.Contracts[contract].Traders(nil, common.HexToAddress(account))
	if err != nil {
		gl.OutLogger.Error("Get account's position data from blockchain error. %v", err)
		return
	}
	if h, err := gl.GetCurrentBlockNumber(); err != nil {
		gl.OutLogger.Error("Get block number error. %v", err)
	} else {
		block = h
	}
	user := model.User{
		Account:   account,
		Margin:    trader.Margin.Int64(),
		Lposition: trader.LongAmount.Uint64(),
		Lprice:    trader.LongPrice.Uint64(),
		Sposition: trader.ShortAmount.Uint64(),
		Sprice:    trader.ShortPrice.Uint64(),
		Block:     block,
	}
	if err := model.UpdateUser(contract, &user); err != nil {
		gl.OutLogger.Error("Update account's data in database error. %s", err.Error())
	}

	if config.Service&0x2 > 0 {
		//update the explosive list
		expUserList[contract].Update(&user)
		//update the takeinterest list
		interestUserList[contract].update(&user)
	}
}
