// main.go

package main

import (
    "fmt"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Define a new counter metric
    counter := prometheus.NewCounter(prometheus.CounterOpts{
        Name: "my_counter",
        Help: "This is my custom counter metric",
    })

    // Register the counter metric with the default Prometheus registry
    prometheus.MustRegister(counter)

    // Increment the counter by 1
    counter.Inc()

    // Set up an HTTP server to expose the registered metrics
    http.Handle("/metrics", promhttp.Handler())
    fmt.Println("Starting HTTP server on port 8080")
    http.ListenAndServe(":8080", nil)
}
