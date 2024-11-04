package main

import (
	"fmt"
	"github.com/esperer/redisperf/api"
	"github.com/esperer/redisperf/redis"
	"github.com/redis/go-redis/v9"
	"log"
)

var redisClient *redis.Client

func main() {

	api.Run()
	config, err := redisconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	redisClient = redisconfig.ConnectRedis(config)
	fmt.Printf("Server Up Redis Address: %s", config.Address)
}
