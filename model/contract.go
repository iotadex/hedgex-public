package model

import "hedgex-server/gl"

//TradePair trade pair model in the database
type TradePair struct {
	Contract      string `json:"contract"`
	MarginAddress string `json:"margin_address"`
	MarginCoin    string `json:"margin_coin"`
	TradeCoin     string `json:"trade_coin"`
	DayOpenPrice  int64  `json:"open_price"`
	IndexPrice    int64  `json:"index_price"`
}

//GetPairs get trade pairs from database
func GetPairs() ([]TradePair, error) {
	rows, err := gl.DB.Query("select contract,margin_address,margin_coin,trade_coin from trade_pair")
	if err != nil {
		return nil, err
	}
	data := make([]TradePair, 0)
	for rows.Next() {
		tp := TradePair{}
		rows.Scan(&tp.Contract, &tp.MarginAddress, &tp.MarginCoin, &tp.TradeCoin)
		data = append(data, tp)
	}
	return data, nil
}
