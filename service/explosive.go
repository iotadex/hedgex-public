package service

import (
	"context"
	"hedgex-server/config"
	"hedgex-server/gl"
	"hedgex-server/model"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var expUserList map[string]*ExplosiveList           //current accounts waiting for be detected to explosive
var explosivedAccounts map[string]*ExplosiveReCheck //have been explosived accounts, used this to verify

func init() {
	expUserList = make(map[string]*ExplosiveList)
	for i := range config.Contract {
		expUserList[config.Contract[i].Address] = NewExplosiveList()
	}
	explosivedAccounts = make(map[string]*ExplosiveReCheck)
}

//StartExplosiveDetectServer, no blocking function
func StartExplosiveDetectServer() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()
	timer := time.NewTicker(config.ExplosiveTick * time.Second)
	for {
		select {
		case <-timer.C:
			auth, err := getAccountAuth()
			if err != nil {
				gl.OutLogger.Error("Get auth error. %v", err)
				continue
			}
			for _, contract := range config.Contract {
				//get the current price of contract
				price, err := Contracts[contract.Address].GetLatestPrice(nil)
				if err != nil {
					gl.OutLogger.Error("Get price from contract error. ", err)
					continue
				}

				node := expUserList[contract.Address].LHead.Next
				for node != nil {
					node = explosive(auth, contract.Address, node, price.Int64(), 1)
				}
				node = expUserList[contract.Address].SHead.Next
				for node != nil {
					node = explosive(auth, contract.Address, node, price.Int64(), -1)
				}
				time.Sleep(time.Second)
			}
		case <-QuitExplosiveDetect:
			return
		}
	}
}

func explosive(auth *bind.TransactOpts, contract string, node *UserNode, price int64, d int64) *UserNode {
	if (node.ExPrice-price)*d > 0 {
		return nil
	}
	nonce, err := EthHttpsClient.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		gl.OutLogger.Error("get nonce error address(%s). %v", publicAddress, err)
		return nil
	}
	auth.Nonce = big.NewInt(int64(nonce))
	if _, err := Contracts[contract].Explosive(auth, common.HexToAddress(node.Account), common.HexToAddress(config.ExplosiveTo)); err != nil {
		gl.OutLogger.Error("Transaction with explosive error. %v", err)
		return nil
	}
	expUserList[contract].Delete(node.Account)
	explosivedAccounts[contract].insert(node)
	return node.Next
}

func StartExplosiveReCheck() {
	ServiceWaitGroup.Add(1)
	defer ServiceWaitGroup.Done()
	timer := time.NewTicker(config.ExplosiveTick * time.Second * 5)
	for {
		select {
		case <-timer.C:
			go func() {
				for _, contract := range config.Contract {
					explosivedAccounts[contract.Address].check(contract.Address)
				}
			}()
		case <-QuitExplosiveReCheck:
			return
		}
	}
}

type UserNode struct {
	model.User
	ExPrice int64 // user's explosive price
	Pre     *UserNode
	Next    *UserNode
}

type ExplosiveList struct {
	LHead *UserNode            // long position user
	SHead *UserNode            // short position user
	Index map[string]*UserNode // the user node's index
	mu    sync.Mutex
}

func NewExplosiveList() *ExplosiveList {
	return &ExplosiveList{
		LHead: &UserNode{},
		SHead: &UserNode{},
		Index: make(map[string]*UserNode),
	}
}

func (el *ExplosiveList) Insert(u *model.User) {
	if (u == nil) || (u.Lposition == u.Sposition) {
		return
	}
	el.mu.Lock()
	defer el.mu.Unlock()
	if _, exist := el.Index[u.Account]; exist {
		return
	}
	keepMargin := (u.Lposition*u.Lprice + u.Sposition*u.Sprice) / 30
	ePrice := (int64(keepMargin) - u.Margin + int64(u.Lposition*u.Lprice) - int64(u.Sposition*u.Sprice)) / (int64(u.Lposition) - int64(u.Sposition))
	var currNode *UserNode
	node := &UserNode{
		User:    *u,
		ExPrice: ePrice,
	}
	if u.Lposition > u.Sposition {
		// find the first node that ExPrice < ePrice
		currNode = el.LHead
		for {
			if (currNode.Next.ExPrice < ePrice) || (currNode.Next == nil) {
				break
			}
			currNode = currNode.Next
		}

	} else {
		currNode = el.SHead
		for {
			if (currNode.Next.ExPrice > ePrice) || (currNode.Next == nil) {
				break
			}
			currNode = currNode.Next
		}
	}
	node.Pre = currNode
	if currNode.Next != nil {
		node.Next = currNode.Next
		currNode.Next.Pre = node
	}
	currNode.Next = node
	el.Index[u.Account] = node
}

func (el *ExplosiveList) Delete(account string) {
	el.mu.Lock()
	defer el.mu.Unlock()
	node, exist := el.Index[account]
	if !exist {
		return
	}
	node.Next.Pre = node.Pre
	node.Pre.Next = node.Next
	delete(el.Index, account)
}

func (el *ExplosiveList) Update(u *model.User) {
	if u == nil {
		return
	}
	el.Delete(u.Account)
	el.Insert(u)
}

type ExplosiveReCheck struct {
	head *UserNode
	tail *UserNode
	mu   sync.Mutex
}

func (erc *ExplosiveReCheck) check(contract string) {
	for erc.head != nil {
		erc.mu.Lock()
		node := erc.head
		if trader, err := Contracts[contract].Traders(nil, common.HexToAddress(node.Account)); err != nil {
			gl.OutLogger.Error("Get account's position data from blockchain error. %s", err.Error())
		} else {
			user := model.User{
				Account:   node.Account,
				Margin:    trader.Margin.Int64(),
				Lposition: trader.LongAmount.Uint64(),
				Lprice:    trader.LongPrice.Uint64(),
				Sposition: trader.ShortAmount.Uint64(),
				Sprice:    trader.ShortPrice.Uint64(),
				Block:     0,
			}
			erc.head = node.Next
			expUserList[contract].Insert(&user)
		}
		erc.mu.Unlock()
	}
}

func (erc *ExplosiveReCheck) insert(node *UserNode) {
	erc.mu.Lock()
	if erc.head == nil {
		erc.head = node
		erc.tail = node
	} else {
		erc.tail.Next = node
		erc.tail = node
	}
	erc.mu.Unlock()
}
