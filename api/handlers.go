package api

import (
	"context"
	"net/http"

	"github.com/ab0626/task-queue-system/metrics"
	"github.com/ab0626/task-queue-system/queue"

	"github.com/gin-gonic/gin"
)

type TaskRequest struct {
	TaskType string `json:"task_type"`
	Payload  string `json:"payload"`
}

func SubmitTaskHandler(c *gin.Context) {
	var req TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := queue.EnqueueTask(context.Background(), req.TaskType, req.Payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enqueue task"})
		return
	}

	metrics.TasksSubmitted.Inc() // Increment the task submitted counter

	c.JSON(http.StatusOK, gin.H{"status": "Task submitted"})
}
