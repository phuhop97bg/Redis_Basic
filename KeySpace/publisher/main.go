package main

import (
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// set key và giá trị
	err := client.Set("key", "value", time.Second*10).Err()
	if err != nil {
		panic(err)
	}
}
