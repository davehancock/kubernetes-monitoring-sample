package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const addr = ":8080"

func main() {

	// Define a custom counter to represent some arbitrary business metric
	trafficCounter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "hits",
			Help: "A metric to represent some endpoint call",
		})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/boom", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Hit! " + time.Now().String())
		trafficCounter.Inc()

		w.Write([]byte("Hello!"))
	})

	// Register the custom counter defined and used above
	prometheus.MustRegister(trafficCounter)

	// Expose all metrics over http for prom, as well as registering some out of the box defaults
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Starting app on address " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
