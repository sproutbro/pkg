package dbx

import (
	"github.com/redis/go-redis/v9"
	"github.com/sproutbro/pkg/config"
)

func NewRedies(rdb *config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     rdb.ADDR,
		Password: rdb.PASS,
		DB:       rdb.DB,
		Protocol: rdb.PROTOCOL,
	})
}
