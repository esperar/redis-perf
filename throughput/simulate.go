package throughput

import (
	"fmt"
	"log"
	"time"
)

func PrintThroughputResults(config *Config) (int, time.Duration) {
	requestCount := handler.Throughput
	if requestCount <= 0 {
		log.Println("Throughput must be greater than 0")
		return 0, 0
	}

	individualTimes := make([]time.Duration, requestCount)
	start := time.Now()

	for i := 0; i < requestCount; i++ {
		reqStart := time.Now()
		// 요청 처리 모의 (간단히 100ms 대기)
		time.Sleep(100 * time.Millisecond)
		reqEnd := time.Since(reqStart)

		// 각 요청 완료 시간 저장
		individualTimes[i] = reqEnd
		fmt.Printf("Request %d completed in: %v\n", i+1, reqEnd)
	}

	totalDuration := time.Since(start)
	fmt.Println("\n--- Throughput Test Results ---")
	fmt.Printf("Total Requests: %d\n", requestCount)
	fmt.Printf("Total Duration: %v\n", totalDuration)
	fmt.Println("--- Individual Request Times ---")

	for i, duration := range individualTimes {
		fmt.Printf("Request %d: %v\n", i+1, duration)
	}
	fmt.Println("--------------------------------")

	// 총 요청 수와 전체 소요 시간 반환
	return requestCount, totalDuration
}
