package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TasksProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_processed_total",
		Help: "Number of tasks processed",
	})

	TasksFailed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_failed_total",
		Help: "Number of tasks that failed processing",
	})

	TasksRetried = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_retried_total",
		Help: "Number of task retries",
	})

	TasksSubmitted = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tasks_submitted_total",
		Help: "Number of tasks submitted",
	})
)

func InitMetrics(port string) {
	prometheus.MustRegister(TasksProcessed, TasksFailed, TasksRetried, TasksSubmitted)
	// http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+port, promhttp.Handler()) // export metrics on port 9100
}
