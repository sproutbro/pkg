package pkg

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func NewRedies(envPath string) *redis.Client {
	if err := godotenv.Load(envPath); err != nil {
		fmt.Println(err)
		return nil
	}

	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
		Protocol: 2,
	})
}
