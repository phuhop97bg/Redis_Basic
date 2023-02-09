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
	topic := "topic_dog"
	pubsub := client.Subscribe(topic)
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
