package test

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Test TestConfig `yaml:"test"`
}

type TestConfig struct {
	Throughput int  `yaml:"throughput"`
	TTL        int  `yaml:"ttl"`
	Failover   bool `yaml:"failover"`
}

func LoadConfig() (*TestConfig, error) {
	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml os: %w", err)
	}
	var config TestConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("error reading yaml unmarshal: %w", err)
	}

	return &config, nil
}
