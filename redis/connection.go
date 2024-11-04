package redisconfig

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"time"
)

var redisClient *redis.Client
var ctx = context.Background()

type RedisConfig struct {
	Address         string `yaml:"address"`
	Password        string `yaml:"password"`
	DB              int    `yaml:"db"`
	MaxMemoryPolicy string `yaml:"maxMemoryPolicy"`
}

func LoadConfig() (*RedisConfig, error) {
	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, err := os.ReadFile(filename)
	var config RedisConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func ConnectRedis(config *RedisConfig) *redis.Client {
	fmt.Printf("Redis Connection: %s", config.Address)
	client := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")
	return client
}

func GetRedisGateWay() *RedisGatewayImpl {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	redisClient = ConnectRedis(config)
	return RedisGatewayImpl{}.New(redisClient, time.Second*5)
}
