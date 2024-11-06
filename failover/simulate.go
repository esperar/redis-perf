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

	// failover 발동
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
