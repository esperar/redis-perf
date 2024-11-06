package failover

import (
	"fmt"
	redisconfig "github.com/esperer/redisperf/redis"
	"gopkg.in/yaml.v3"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type FailoverTestResult struct {
	ExpectedCount string        `json:"expected_count"`
	ResultCount   string        `json:"result_count"`
	ErrorMargin   string        `json:"error_margin"`
	FailoverTime  time.Duration `json:"failover_time"`
}

type RedisNodeConfig struct {
	Master string `json:"master"`
}

func loadConfig() (*RedisNodeConfig, error) {
	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, err := os.ReadFile(filename)
	var config RedisNodeConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

var redisGateway = redisconfig.GetRedisGateWay()

func PrintFailoverTestResult() (*FailoverTestResult, error) {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Fail to load redis master node config")
		return nil, err
	}
	// insert test data
	dataCount := 10000
	for i := 0; i < dataCount; i++ {
		err := redisGateway.SetData(fmt.Sprintf("key%d", i), strconv.Itoa(i))
		if err != nil {
			return nil, fmt.Errorf("failed to set data: %d", i)
		}
	}

	startTime := time.Now()

	// Redis Sentinel을 사용해 강제로 페일오버 트리거하는 스크립트를 발동시킴
	// 예제 명령어는 Redis Sentinel이 설정된 환경에서 마스터 노드를 수동으로 페일오버시킴.
	// 여러가지가 있지만 이 예제에서는 Redis Sentinel의 `SENTINEL failover <master-name>` 사용

	/*
	   Example Redis Sentinel failover command:
	   $ redis-cli -p <sentinel_port> SENTINEL failover <master-name>
	   $ redis-cli -p 26379 SENTINEL failover mymaster
	*/

	sentinelPort := "26379" // Replace with your actual sentinel port
	master := config.Master // Replace with your actual master name

	cmd := exec.Command("redis-cli", "-p", sentinelPort, "SENTINEL", "failover", master)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to execute failover command: %v", err)
	}

	// failover 완료 시간을 5초로 가정 -> 그러나 페일오버가 완료되었을 신호를 받고 실행하는게 더 좋아보임
	// 5초정도 기다리고 만약 페일오버가 완료되지 않았다면 5초정도 더 기다리고 만약 그 이후에도 그대로라면 실패처리
	time.Sleep(5 * time.Second)

	// 페일오버 후 데이터 검증
	resultCount := 0
	for i := 0; i < dataCount; i++ {
		val, err := redisGateway.GetData(fmt.Sprintf("key%d", i))
		if err == nil && val == fmt.Sprintf("%d", i) {
			resultCount++
		}
	}

	failoverTime := time.Since(startTime)
	errorMargin := dataCount - resultCount

	fmt.Println("========== Failover Test Result ==========")
	fmt.Println("---- Test Summary ----")
	fmt.Printf("Expected Count : %d\n", dataCount)
	fmt.Printf("Result Count   : %d\n", resultCount)
	fmt.Printf("Error Margin   : %d\n", errorMargin)
	fmt.Printf("Failover Time  : %v\n", failoverTime)
	fmt.Println("==========================================")

	result := &FailoverTestResult{
		ExpectedCount: strconv.Itoa(dataCount),
		ResultCount:   strconv.Itoa(resultCount),
		ErrorMargin:   strconv.Itoa(errorMargin),
		FailoverTime:  failoverTime,
	}

	return result, nil
}
