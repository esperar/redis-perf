package api

import (
	"encoding/json"
	"github.com/esperer/redisperf/redis"
	"net/http"
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

func HandleSimulateFailover(w http.ResponseWriter, r *http.Request) {

}

func HandleSimulateThroughput(w http.ResponseWriter, r *http.Request) {

}

func HandleSimulateTTL(w http.ResponseWriter, r *http.Request) {

}
