package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(redisHost string, redisPort string) {
	fmt.Println("[INFO] Inside InitRedis")
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
	})

	// testing the connection

	if pong := redisClient.Ping(context.Background()); pong.String() != "ping: PONG" {
		fmt.Println("-------------Error connection redis ----------:", pong)
		panic(pong)
	}

	fmt.Println("[SUCCESS] Redis: Connection established")
}
