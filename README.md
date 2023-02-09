# Redis_Basic

"go mod download"

I, Redis Pub/Sub
a. Run Subscriber  "go run .\PubSub\subscriber\main.go"
b. Run Publisher "go run .\PubSub\publisher\main.go"
II.1, Redis keyspace notifications: __keyspace@0__
a. Run "go run .\KeySpace\subscriber_event\main.go"
b. Go to RedisCLI to set, setex, del, expire...
II.2, Redis keyevent notifications: __keyevent@0__
a. Run "go run .\KeySpace\subscriber_space\main.go"
b. Go to RedisCLI to set, setex, del, expire...
