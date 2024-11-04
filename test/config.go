package test

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type TestConfig struct {
	Throughput int  `yaml:throughput`
	TTL        int  `yaml:ttl`
	Failover   bool `yaml:failover`
}

func LoadConfig() (*TestConfig, error) {
	filename, _ := filepath.Abs("config/config.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	var config TestConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
