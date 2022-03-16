package kline

import (
	"context"
	"hedgex-public/config"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

type RedisKline struct {
	tradePair string
	c         *redis.Client
}

func NewRedisKline(tp string) *RedisKline {
	return &RedisKline{
		tradePair: tp,
		c: redis.NewClient(&redis.Options{
			Addr:     config.Redis.Addr,
			Password: "",
			DB:       0,
		}),
	}
}

func (rk *RedisKline) Get(kt string, count int) ([][5]int64, error) {
	vs, err := rk.c.LRange(context.Background(), rk.tradePair+kt, 0, int64(count-1)).Result()
	if err != nil {
		return nil, err
	}

	count = len(vs)
	candles := make([][5]int64, count)
	for i, v := range vs {
		datas := strings.Split(v, ".")
		var candle [5]int64
		for j := range datas {
			candle[j], _ = strconv.ParseInt(datas[j], 10, 64)
		}
		candles[count-1-i] = candle
	}

	return candles, nil
}

func (rk *RedisKline) GetCurrent(kt string) ([5]int64, error) {
	v, err := rk.c.LIndex(context.Background(), rk.tradePair+kt, 0).Result()
	if err != nil {
		if err == redis.Nil {
			return [5]int64{}, nil
		}
		return [5]int64{}, err
	}

	datas := strings.Split(v, ".")
	var candle [5]int64
	for i := range datas {
		if candle[i], err = strconv.ParseInt(datas[i], 10, 64); err != nil {
			break
		}
	}
	return candle, err
}

func (rk *RedisKline) Append(kt string, currentData [5]int64) error {
	key := rk.tradePair + kt
	str := strconv.FormatInt(currentData[0], 10)
	for i := 1; i < 5; i++ {
		str += "." + strconv.FormatInt(currentData[i], 10)
	}

	candle, _ := rk.GetCurrent(kt)
	if candle[4] == currentData[4] {
		return rk.c.LSet(context.Background(), key, 0, str).Err()
	}

	if err := rk.c.LPush(context.Background(), key, str).Err(); err != nil {
		return err
	}

	count, err := rk.c.LLen(context.Background(), key).Result()
	if err != nil {
		return err
	}
	for i := 0; i < int(count)-config.MaxKlineCount; i++ {
		err := rk.c.RPop(context.Background(), key).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
