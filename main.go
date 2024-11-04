package redisperf

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
