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
	topic := "topic_dog*"
	pubsub := client.PSubscribe(topic)
	defer pubsub.Close()

	fmt.Printf("subcribe successfully topic %s \n", topic)
	// Đăng ký với chủ đề
	_, err := pubsub.Receive()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Nhận tin nhắn từ chủ đề
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, "------------------", msg.Payload)
	}
}
