package main

import (
	"flag"
	"fmt"

	// "os"
	"time"

	"github.com/ab0626/task-queue-system/api"
	"github.com/ab0626/task-queue-system/metrics" // Importing the metrics package
	"github.com/gin-gonic/gin"
)

func main() {
	mode := flag.String("mode", "api", "Mode to run: api or worker")
	flag.Parse()

	metricsPort := "9100"
	if *mode == "worker" {
		metricsPort = "9101"
	}

	metrics.InitMetrics(metricsPort) // Initialize Prometheus metrics (removed: undefined function)

	if *mode == "worker" {
		fmt.Println("Running in worker mode")
		// Example worker loop (replace with real queue logic)
		for {
			taskType := "exampleType"
			payload := "examplePayload"
			if err := processTask(taskType, payload); err != nil {
				metrics.TasksFailed.Inc()
				metrics.TasksRetried.Inc()
			} else {
				metrics.TasksProcessed.Inc()
			}
			time.Sleep(5 * time.Second)
		}
	}

	r := gin.Default()
	r.POST("/submit", api.SubmitTaskHandler)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running"})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running"})
	})
	r.Run(":8080") // API server runs on port 8080
}

// processTask is a placeholder implementation for demonstration purposes.
func processTask(taskType string, payload string) error {
	_ = taskType
	_ = payload
	// Simulate successful processing; return nil or an error as needed.
	return nil
}
