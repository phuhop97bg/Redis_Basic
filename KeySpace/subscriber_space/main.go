package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := client.ConfigSet("notify-keyspace-events", "KEA").Err(); err != nil {
		panic(err)
	}

	channel := "__keyspace@*__:key12*"
	pubsubSpace := client.PSubscribe(channel)
	defer pubsubSpace.Close()
	fmt.Printf("subscribe ok : channel %s \n", channel)
	for {
		msg, err := pubsubSpace.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, "-------", msg.Payload)
	}
}
