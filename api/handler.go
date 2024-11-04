package api

import (
	"encoding/json"
	"fmt"
	"github.com/esperer/redisperf/redis"
	"github.com/esperer/redisperf/test"
	"github.com/esperer/redisperf/throughput"
	"log"
	"net/http"
	"strconv"
)

var redisGateway = redisconfig.GetRedisGateWay()

func OkHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "OK"}
	json.NewEncoder(w).Encode(response)
}

func HandleHealthCheckRedisHandler(w http.ResponseWriter, r *http.Request) {
	err := redisGateway.Ping(r.Context())
	if err != nil {
		response := map[string]string{"message": "PONG Failed"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "PONG"}
	json.NewEncoder(w).Encode(response)
}

func HandleSimulateThroughput(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("원욱이")
	config, err := test.LoadConfig()
	if err != nil {
		response := map[string]string{"message": "Throughput Test Failed Because Can'not load Test Config"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		log.Println(err)
		return
	}
	requestCount, averageDuration, totalDuration := throughput.PrintThroughputResults(config)
	response := map[string]string{
		"requestCount":    strconv.Itoa(requestCount) + "req",
		"averageDuration": averageDuration.String(),
		"totalDuration":   totalDuration.String(),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func HandleSimulateFailover(w http.ResponseWriter, r *http.Request) {

}

func HandleSimulateTTL(w http.ResponseWriter, r *http.Request) {

}
