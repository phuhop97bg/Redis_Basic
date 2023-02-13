package main

import (
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.8.12.195:7002",
		DB:       15,
		Password: "p8eruwyODJ5Vnr2w3qlc",
	})

	// set key và giá trị
	key := "key1234"
	value := "value1231424125"
	err := client.Set(key, value, time.Second*10).Err()
	if err != nil {
		panic(err)
	}
	client.LPush("myset123", key)
}
