package opc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	opcReadsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "opc_reads_total",
			Help: "Counts the total number of OPC tags read.",
		},
		[]string{"status"}, // "success" == 0 or "failed" == 1
	)

	opcReadsDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "opc_reads_duration_seconds",
			Help:    "Read duration in seconds from OPC server.",
			Buckets: prometheus.ExponentialBuckets(0.000001, 10, 6), // start with 500 ns, add 500 ns for 5 buckets.
		},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(opcReadsCounter)
	prometheus.MustRegister(opcReadsDuration)
}

func StartMonitoring(port string) {
	var p string
	if port == "" {
		p = ":8080"
	} else {
		p = port
	}
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Fatal(http.ListenAndServe(p, nil))
	}()
}