package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.117.58.9:8379",
		Password: "i3Fp9t1Rbcw",
		DB:       0,
	})

	// ZSet
	//seconds := float64(time.Now().Unix())
	//fmt.Println("time now seconds: ", time.Now().Unix())
	//fmt.Println("time now float64: ", seconds)
	//
	//if err := rdb.ZAdd(ctx, "myset", redis.Z{
	//	Score:  seconds,
	//	Member: 3,
	//}).Err(); err != nil {
	//	return
	//}
	//
	//members, err := rdb.ZRangeByScore(ctx, "myset", &redis.ZRangeBy{
	//	Min:   "-inf",
	//	Max:   fmt.Sprint(time.Now().Unix() - 10),
	//	Count: 100,
	//}).Result()
	//if err != nil {
	//	return
	//}
	//for _, member := range members {
	//	fmt.Println(member)
	//}
	//
	//rem := rdb.ZRem(ctx, "myset", 5577006791947779410)
	//fmt.Println(rem)

	// List
	if err := rdb.RPush(ctx, "mylist", "hello").Err(); err != nil {
		return
	}

	if err := rdb.RPush(ctx, "mylist", "hello2").Err(); err != nil {
		return
	}

	for {
		result, err := rdb.LPop(ctx, "mylist").Result()
		if err != nil {
			return
		}

		fmt.Println(result)

		if result == "" {
			break
		}
	}
}
