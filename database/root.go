package database

import (
	"context"
	"github.com/go-redis/redis/v7"
)

var (
	Ctx = context.Background()
)

func RedisInit() *redis.Client {
	opt, _ := redis.ParseURL()
	client := redis.NewClient(opt)

	return client
}

func MySql()
