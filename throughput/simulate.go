package throughput

import (
	"fmt"
	redisconfig "github.com/esperer/redisperf/redis"
	"github.com/esperer/redisperf/test"
	"log"
	"time"
)

var redisGateway = redisconfig.GetRedisGateWay()

func PrintThroughputResults(config *test.TestConfig) (int, time.Duration, time.Duration) {
	requestCount := config.Throughput
	keyCount := 30
	individualTimes := make([]time.Duration, keyCount)
	for i := 1; i <= keyCount; i++ {
		key := fmt.Sprintf("test_key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		err := redisGateway.SetData(key, value)
		if err != nil {
			log.Printf("Failed to set key %s: %v\n", key, err)
		}
	}
	// Measure time taken for each GET request
	for i := 1; i <= requestCount; i++ {
		key := fmt.Sprintf("test_key_%d", (i%keyCount)+1)
		reqStart := time.Now()
		_, err := redisGateway.GetData(key)
		reqEnd := time.Since(reqStart)
		if err != nil {
			log.Printf("Failed to get key %s: %v\n", key, err)
		}
		individualTimes[i-1] = reqEnd
		fmt.Printf("Request %d for key %s completed in: %v\n", i, key, reqEnd)
	}
	// Calculate total and average duration
	var totalDuration time.Duration
	for _, duration := range individualTimes {
		totalDuration += duration
	}
	averageDuration := totalDuration / time.Duration(keyCount)
	fmt.Println("\n--- Throughput Test Results ---")
	fmt.Printf("Total Keys: %d\n", keyCount)
	fmt.Printf("Total Duration: %v\n", totalDuration)
	fmt.Printf("Average Request Duration: %v\n", averageDuration)
	fmt.Println("--------------------------------")
	return requestCount, averageDuration, totalDuration
}
