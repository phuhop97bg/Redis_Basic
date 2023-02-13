package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.8.12.195:7002",
		DB:       15,
		Password: "p8eruwyODJ5Vnr2w3qlc",
	})
	ctx := context.Background()
	if err := client.ConfigSet(ctx, "notify-keyspace-events", "KEA").Err(); err != nil {
		panic(err)
	}
	channel := "__keyevent@15__:set"
	pubsubEvent := client.PSubscribe(ctx, channel)
	fmt.Printf("subscribe ok : channel %s \n", channel)
	defer pubsubEvent.Close()

	for {
		msg, err := pubsubEvent.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		val, err := client.RPop(ctx, "myset123").Result()
		if val == "" || err != nil {
			continue
		} else {
			fmt.Println("subscriber3:------------", msg.Channel, msg.Payload)
		}
	}
}
