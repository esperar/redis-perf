package redisperf

import (
	"fmt"
	"github.com/esperer/redisperf/api"
	"github.com/esperer/redisperf/redis"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var redisClient *redis.Client

func main() {
	config, err := redisconfig.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	redisClient = redisconfig.ConnectRedis(config)

	http.HandleFunc("/healthredis", api.HealthCheckRedisHandler)

	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
