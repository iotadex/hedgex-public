# Hedgex Single User Interface
## H5 rest api， URL : https://triple.fi/api
### GET /contract/trade_pairs get the trade pair list
#### result
```json
{
    "result": true,
    "data":[
		{
            "contract":"contract address",
            "margin_coin":"margin coin name",
            "trade_coin":"trade coin name",
            "open_price":1234,
            "index_price":3462
        }
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|contract|string|contract address, as "0x0x1660854c03b461E6BC07f94567D1D6E6bF99a1A9"|
|open_price|int64|the open price of current day|
|index_price|int64|the index price of current time|

### Get /contract/pair_params?contract={address} get the trade pair's params
#### return
```json
{
    "result": true,
    "data":{
        "leverage":25,
		"min_amount":0.01,
	    "keep_margin_scale":60,
		"fee_rate":0.0006,
		"single_close_limit_rate":0.03,
		"single_open_limit_rate":0.1,
		"r_open":0.4,
		"r_price":0.3,
		"token0":"0x2E5591820Dcd82Bf75B369665Ca81eA2Fe54BfB5",
		"token0_decimal":1000000,
		"daily_interest_rate_base":0.001
    }
}
```
|Name|Type|Description|
|---|:--:|---|
|leverage|int|leverage|
|min_amount|float64|min_amount = 10^amountDecimal|
|keep_margin_scale|int|keepMarginScale/divConst|
|fee_rate|float64|feeRate/divConst|
|single_close_limit_rate|float64|singleCloseLimitRate/divConst|
|single_open_limit_rate|float64|singleOpenLimitRate/divConst|
|r_open|float64|poolNetAmountRateLimitOpen/divConst|
|r_price|float64|poolNetAmountRateLimitPrice/divConst|
|token0|string|token0|
|token0_decimal|int|token0_decimal = token0Decimal = 10^decimals|
|daily_interest_rate_base|float64|dailyInterestRateBase/divConst|

### GET /contract/kline get the kline data from current time(now)
#### parameter
|Name|Type|Description|
|---|:--:|---|
|contract|string|contract address, as "0x0x1660854c03b461E6BC07f94567D1D6E6bF99a1A9"|
|type|string|kline's type. There are "m1","m5","m10","m15","m30","h1","h2","h4","h6","h12","d1"|
|count|int|data's count, max is 2000|
#### result
```json
{
    "result": true,
    "data":[
		["open","high","low","close","ts"]
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|open|int|open price|
|high|int|highest price|
|low|int|lowest price|
|close|int|close price|
|ts|int|the timestamp as second, the earliest time is the first one|


### GET /contract/position get the traders' count for every contract
#### result
```json
{
    "result": true,
    "data":{
        "contract_address":{
            "long":2,
            "short":3,
            "total":4,
        }
    }
}
```

### GET /contract/trade?contract={address}&count=30 get the contract's trade
```json
{
    "result": true,
    "data":[
		{
            "direction":1,
            "amount":10,
            "price":52346,
            "block":12234,
            "ts":"2022-03-02 23:29:37"
        }
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|block|int|the blocknumber when explosive happen|
|direction|int|1:open long,-1:open short,-2:close long,2:close short|
|amount|int|piece|
|price|int|the price of one piece|

### GET /contract/explosive?contract={address}&count=30 get the contract's trade
```json
{
    "result": true,
    "data":[
		{
            "direction":1,
            "amount":10,
            "price":52346,
            "block":12234,
            "ts":"2022-03-02 23:29:37"
        }
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|block|int|the blocknumber when explosive happen|
|direction|int|1:open long,-1:open short,-2:close long,2:close short|
|amount|int|piece|
|price|int|the price of one piece|

### GET /account?contract={address} get all the trader's detail data by contract
#### result
```json
{
    "result": true,
    "data":[
		{
            "account":"trader address",
            "margin":1,
            "lposition":0,
            "lprice":0,
            "sposition":0,
            "sprice":0,
            "interestDay":1234,
            "block":123445
        }
	]
}
```

### GET /account/trade get user's trade list
#### parameter
|Name|Type|Description|
|---|:--:|---|
|contract|string|contract address, as "0x0x1660854c03b461E6BC07f94567D1D6E6bF99a1A9"|
|account|string|user's wallet address|
|count|int|the count of list, max is 200 latest|
#### result
```json
{
    "result": true,
    "data":[
		{
            "tx":"transaction hash",
            "block":12234,
            "direction":1,
            "amount":10,
            "price":52346
        }
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|block|int|the blocknumber when explosive happen|
|direction|int|1:open long,-1:open short,-2:close long,2:close short|
|amount|int|piece|
|price|int|the price of one piece|

### GET /account/explosive get the user's explosive list
#### parameter
|Name|Type|Description|
|---|:--:|---|
|contract|string|contract address, as "0x0x1660854c03b461E6BC07f94567D1D6E6bF99a1A9"|
|account|string|user's wallet address|
|count|int|the count of list, max is 100 latest|
#### result
```json
{
    "result": true,
    "data":[
		{
            "tx":"transaction hash",
            "block":12234,
            "amount":10,
            "price":52346,
            "direction":1
        }
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|block|int|the blocknumber when explosive happen|
|amount|int|piece|
|price|int|the price of one piece|
|direction|int| as same with trade|

### GET /account/interest get user's interest records
#### parameter
|Name|Type|Description|
|---|:--:|---|
|contract|string|contract address, as "0x0x1660854c03b461E6BC07f94567D1D6E6bF99a1A9"|
|account|string|user's wallet address|
|count|int|the count of list, max is 100 latest|
#### result
```json
{
    "result": true,
    "data":[
		{
            "tx":"transaction hash",
            "block":12234,
            "amount":10,
            "price":52346
        }
	]
}
```
|Name|Type|Description|
|---|:--:|---|
|block|int|the blocknumber when explosive happen|
|amount|int|the margin's amount|
|price|int|the price of one piece|

### GET /odds/testcoin?account={address} Send testcoin to the account, each account can get 10000usdp.
#### result
```json
{
    "result": true,
    "data":""
}
```

### GET /odds/add_email?email={address} send user's email to us
#### result
```json
{
    "result": true,
    "data":""
}
```

### GET /odds/emails?from={timestamp}&to={timestamp} get email addresses with timestamp
#### result
```json
{
    "result": true,
    "data":["email1","email2"]
}
```

## kline websockt， URL : wss://triple.fi/wss/kline
```
When the websocket connection is established. Send the kline's type string with contract address(address:m1,address:m5,...), then the kline data will be send from the server.
The data you will be received is a array [5]int64 : [open, high, low, close, ts]
```

# Blockchain Network And Coins' Address

## eth-mainnet
|coin|address|
|---|:--:|
|weth|0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2|
|usdt|0xdac17f958d2ee523a2206206994597c13d831ec7|
|usdc|0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48|
|bnb |0xB8c77482e45F1F44dE1745F52C74426C631bDD52|

## bsc-mainnet
|coin|address|
|---|:--:|
|weth|0x2AaB30A909748945d42c7d29d3CD44A5680caB1D|
|usdt|0x55d398326f99059fF775485246999027B3197955|
|busd|0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56|
|wbnb|0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c|

,"0xB6D0FBe198bf48c7E7dad8b457994a0dBac795Ef"

39e10402beff72d338b4c16b5094f88c94330aa32a17351bf5be05da92671a4d