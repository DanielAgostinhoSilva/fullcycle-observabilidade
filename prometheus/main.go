package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
)

var onlineUsers = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "goapp_online_users",
	Help: "Online users",
	ConstLabels: map[string]string{
		"course": "fullcycle",
	},
})

var httpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "goapp_http_requests_total",
	Help: "Count of all HTTP request for goapp",
}, []string{})

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(onlineUsers)
	r.MustRegister(httpRequestTotal)

	go func() {
		for {
			onlineUsers.Set(float64(rand.Intn(2000)))
		}
	}()

	home := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Full Cycle"))
	})

	http.Handle("/", promhttp.InstrumentHandlerCounter(httpRequestTotal, home))

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8181", nil))
}
