package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const addr = ":8080"

func main() {

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit! " + time.Now().String())
		w.Write([]byte("Hello!"))
	})

	fmt.Println("Starting app on address " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
