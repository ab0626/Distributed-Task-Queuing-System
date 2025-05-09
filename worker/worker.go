package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "redis:6379", // match service name in docker-compose
})

// Simulated task processing
func processTask(taskType, payload string) error {
	fmt.Printf("Processing Task - Type: %s | Payload: %s\n", taskType, payload)
	// simulate task failure 30% of the time
	if time.Now().UnixNano()%3 == 0 {
		return fmt.Errorf("simulated task failure")
	}
	return nil
}

func main() {
	for {
		task, err := rdb.BRPop(ctx, 0*time.Second, "task_queue").Result()
		if err != nil {
			log.Println("Error fetching task:", err)
			continue
		}

		parts := strings.SplitN(task[1], ":", 2)
		if len(parts) != 2 {
			log.Println("Invalid task format")
			continue
		}

		taskType := parts[0]
		payload := parts[1]
		err = processTask(taskType, payload)

		if err != nil {
			fmt.Println("Task failed. Retrying or moving to DLQ...")
			retryKey := fmt.Sprintf("retry:%s", task[1])
			retries, _ := rdb.Incr(ctx, retryKey).Result()

			if retries > 3 {
				_ = rdb.LPush(ctx, "dead_letter_queue", task[1]).Err()
				rdb.Del(ctx, retryKey) // clear retry counter
				fmt.Println("Moved to DLQ:", task[1])
			} else {
				_ = rdb.LPush(ctx, "task_queue", task[1]).Err()
				fmt.Printf("Retry %d for task: %s\n", retries, task[1])
			}
		}
	}
}
