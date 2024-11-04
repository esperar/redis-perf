package redisconfig

import (
	fmt "fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

var redisClient *redis.Client

type RedisConfig struct {
	Address         string `yaml:address`
	Password        string `yaml:password`
	DB              int    `yaml:db`
	MaxMemoryPolicy string `yaml:maxMemoryPolicy`
}

func LoadConfig() (*RedisConfig, error) {
	filename, _ := filepath.Abs("config/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config RedisConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func ConnectRedis(config *RedisConfig) *redis.Client {
	addr := fmt.Sprintf("Redis Connection: %s", config.Address)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	})

	fmt.Println("Connected to Redis")
	return client
}
