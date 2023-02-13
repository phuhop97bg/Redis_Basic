package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.8.12.195:7002",
		DB:       14,
		Password: "p8eruwyODJ5Vnr2w3qlc",
	})
	topic := "topic_dog123"
	msg := "ádasd"
	err := client.Publish(topic, msg).Err()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("publish to topic %s ok", topic)
	}

}
