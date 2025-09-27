package msg

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Logger struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Logger {
	return &Logger{rdb: rdb}
}

const onlyErrKey = "only:error:counter"

func (l *Logger) IncrCounter(ctx context.Context, key string) (int64, error) {

	return l.rdb.Incr(ctx, key).Result()
}

func (l *Logger) Hset(ctx context.Context, module, errType, msg string) error {
	l.IncrCounter(ctx, module)
	payload := map[string]any{
		"module": module,
		"type":   errType,
		"msg":    msg,
		"time":   time.Now().UTC().Format(time.RFC3339),
	}
	data, _ := json.Marshal(payload)
	return l.rdb.Publish(ctx, "app:error", data).Err()
}
