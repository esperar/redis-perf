package api

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net/http"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func HealthCheckRedisHandler(w http.ResponseWriter, r *http.Request) {
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		http.Error(w, "Redis connection is down", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Redis is healthy"))
}
