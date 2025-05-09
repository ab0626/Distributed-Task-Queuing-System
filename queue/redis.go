package queue

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "redis:6379", // docker-compose will use 'redis' as hostname
})

func EnqueueTask(ctx context.Context, taskType, payload string) error {
	task := taskType + ":" + payload
	return rdb.LPush(ctx, "task_queue", task).Err()
}
