package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := client.ConfigSet("notify-keyspace-events", "KEA").Err()
	topic := "topic_dog"
	msg := "Hello Dog"
	err = client.Publish(topic, msg).Err()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("publish to topic %s ok", topic)
}
