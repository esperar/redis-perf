package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", OkHandler).Methods("GET")
	router.HandleFunc("/ping", HandleHealthCheckRedisHandler).Methods("GET")
	router.HandleFunc("/failover", HandleSimulateFailover).Methods("GET")
	router.HandleFunc("/throughput", HandleSimulateThroughput).Methods("GET")
	router.HandleFunc("/pipeline", HandlePipeline).Methods("GET")
	router.HandleFunc("/ttl", HandleSimulateTTL).Methods("GET")

	fmt.Println("Server Start!! :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
