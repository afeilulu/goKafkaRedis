package config

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	// RedisClient ......
	RedisClient *redis.Client
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "Xinlu20180302",
		DB:       0,
	})

	res, err := RedisClient.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Redis " + res)

	// psNewMessage := RedisClient.PSubscribe("__keyevent@*__:expired")

	// go func() {
	// 	defer psNewMessage.Close()
	// 	for {
	// 		msg, err := psNewMessage.ReceiveMessage()
	// 		if err != nil {
	// 			fmt.Printf("kaboooom: %s\n", err.Error())
	// 			continue
	// 		}
	// 		fmt.Printf("Message: %s %s %s\n", msg.Channel, msg.Pattern, msg.Payload)
	// 	}
	// }()

}