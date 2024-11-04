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
	w.WriteHeader(http.StatusOK)
	redisGateway.Ping(r.Context())
	response := map[string]string{"message": "PONG"}
	json.NewEncoder(w).Encode(response)
}
