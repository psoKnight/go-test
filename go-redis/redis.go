package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.117.58.9:8379",
		Password: "i3Fp9t1Rbcw",
		DB:       0,
	})

	seconds := float64(time.Now().Unix())
	fmt.Println("time now seconds: ", time.Now().Unix())

	if err := rdb.ZAdd(ctx, "myset", redis.Z{
		Score:  seconds,
		Member: 3,
	}).Err(); err != nil {
		return
	}

	members, err := rdb.ZRangeByScore(ctx, "myset", &redis.ZRangeBy{
		Min:   "-inf",
		Max:   fmt.Sprint(time.Now().Unix() - 10),
		Count: 100,
	}).Result()
	if err != nil {
		return
	}
	for _, member := range members {
		fmt.Println(member)
	}

	rem := rdb.ZRem(ctx, "myset", 5577006791947779410)
	fmt.Println(rem)
}
