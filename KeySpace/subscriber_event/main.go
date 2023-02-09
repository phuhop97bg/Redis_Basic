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
	channel := "__keyevent@*__:expired"
	pubsubEvent := client.PSubscribe(channel)
	fmt.Printf("subscribe ok : channel %s \n", channel)
	defer pubsubEvent.Close()

	for {
		msg, err := pubsubEvent.ReceiveMessage()
		if err != nil {
			panic(err)
		}

		fmt.Println(msg.Channel, msg.Payload)
	}
}
