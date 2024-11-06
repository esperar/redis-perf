package failover

import (
	"fmt"
	redisconfig "github.com/esperer/redisperf/redis"
	"strconv"
	"time"
)

type FailoverTestResult struct {
	ExpectedCount int           `json:"expected_count"`
	ResultCount   int           `json:"result_count"`
	ErrorMargin   int           `json:"error_margin"`
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

	result := &FailoverTestResult{
		ExpectedCount: dataCount,
		ResultCount:   resultCount,
		ErrorMargin:   errorMargin,
		FailoverTime:  failoverTime,
	}

	return result, nil
}
