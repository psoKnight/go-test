package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type RedisStr struct {
	A string `json:"A"`
	B int    `json:"B"`
}

// MarshalBinary 序列化
func (m *RedisStr) MarshalBinary() (data []byte, err error) {
	fmt.Println("MarshalBinary")
	return json.Marshal(m)
}

// UnmarshalBinary 反序列化
func (m *RedisStr) UnmarshalBinary(data []byte) error {
	fmt.Println("UnmarshalBinary")
	return json.Unmarshal(data, m)

}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.117.58.9:8379",
		Password: "c06pTs8i37S",
		DB:       0,
	})

	// string
	bRedisstr := RedisStr{
		A: "a",
		B: 1,
	}
	if err := rdb.Set(ctx, "string_key_1", &bRedisstr, time.Duration(10)*time.Second).Err(); err != nil {
		fmt.Println(err)
		return
	}

	aRedisstr := RedisStr{}
	if err := rdb.Get(ctx, "string_key_1").Scan(&aRedisstr); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(aRedisstr)
	}

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
	//if err := rdb.RPush(ctx, "mylist", "hello").Err(); err != nil {
	//	return
	//}
	//
	//if err := rdb.RPush(ctx, "mylist", "hello2").Err(); err != nil {
	//	return
	//}
	//
	//for {
	//	result, err := rdb.LPop(ctx, "mylist").Result()
	//	if err != nil {
	//		return
	//	}
	//
	//	fmt.Println(result)
	//
	//	if result == "" {
	//		break
	//	}
	//}
}
