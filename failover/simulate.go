package failover

import (
	"fmt"
	redisconfig "github.com/esperer/redisperf/redis"
	"strconv"
	"time"
)

type FailoverTestResult struct {
	ExpectedCount string        `json:"expected_count"`
	ResultCount   string        `json:"result_count"`
	ErrorMargin   string        `json:"error_margin"`
	FailoverTime  time.Duration `json:"failover_time"`
}

var redisGateway = redisconfig.GetRedisGateWay()

func PrintFailoverTestResult() (*FailoverTestResult, error) {
	// insert test data
	dataCount := 10000
	for i := 0; i < dataCount; i++ {
		err := redisGateway.SetData(fmt.Sprintf("key%d", i), strconv.Itoa(i))
		if err != nil {
			return nil, fmt.Errorf("failed to set data: %d", i)
		}
	}

	startTime := time.Now()

	// Redis Sentinel을 사용해 강제로 페일오버 트리거하는 스크립트
	// 예제 명령어는 Redis Sentinel이 설정된 환경에서 마스터 노드를 수동으로 페일오버시킴.
	// 여러가지가 있지만 이 예제에서는 Redis Sentinel의 `SENTINEL failover <master-name>` 사용
	// master-name은 Sentinel 설정 파일의 마스터 설정으로 config.yml 사용 예정

	/*
	   Example Redis Sentinel failover command:
	   $ redis-cli -p <sentinel_port> SENTINEL failover <master-name>

	   예시:
	   $ redis-cli -p 26379 SENTINEL failover mymaster
	*/

	// TODO 위 스크립트를 cmd Exec을 통해 실행시키는것이 좋을 것으로 보

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
