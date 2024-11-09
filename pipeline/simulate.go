package pipeline

import (
	"fmt"
	"github.com/esperer/redisperf/redis"
	"github.com/esperer/redisperf/test"
	"time"
)

type PipelineResult struct {
	BasicSetTime    time.Duration `json:"basic_set_time"`
	BasicGetTime    time.Duration `json:"basic_get_time"`
	PipelineSetTime time.Duration `json:"pipeline_set_time"`
	PipelineGetTime time.Duration `json:"pipeline_get_time"`
}

var redisGateway = redisconfig.GetRedisGateWay()

func PrintPipelineTestResult(config *test.TestConfig) (*PipelineResult, error) {
	testCount := config.Throughput

	data := make(map[string]string, testCount)
	for i := 0; i < testCount; i++ {
		data[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	keys := make([]string, testCount)
	for i := 0; i < testCount; i++ {
		keys[i] = fmt.Sprintf("key%d", i)
	}

	basicSetTime := measureBasicSet(data)
	basicGetTime := measureBasicGet(keys)
	pipelineSetTime := measurePipelineSet(data)
	pipelineGetTime := measurePipelineGet(keys)

	result := &PipelineResult{
		BasicSetTime:    basicSetTime,
		BasicGetTime:    basicGetTime,
		PipelineSetTime: pipelineSetTime,
		PipelineGetTime: pipelineGetTime,
	}

	return result, nil

}

func measureBasicSet(data map[string]string) time.Duration {
	start := time.Now()
	for key, value := range data {
		err := redisGateway.SetData(key, value)
		if err != nil {
			fmt.Println("Error Setting Data Set: ", err)
			return -1
		}
	}
	return time.Since(start)
}

func measureBasicGet(keys []string) time.Duration {
	start := time.Now()
	for _, key := range keys {
		_, err := redisGateway.GetData(key)
		if err != nil {
			fmt.Println("Error Getting Data", err)
			return -1
		}
	}
	return time.Since(start)
}

func measurePipelineSet(data map[string]string) time.Duration {
	start := time.Now()
	err := redisGateway.SetDataByPipeline(data)
	if err != nil {
		fmt.Println("Error Setting data by pipeline", err)
		return -1
	}
	return time.Since(start)
}

func measurePipelineGet(keys []string) time.Duration {
	start := time.Now()
	_, err := redisGateway.GetDataByPipeline(keys)
	if err != nil {
		fmt.Println("Error Getting data by pipeline", err)
	}
	return time.Since(start)
}
