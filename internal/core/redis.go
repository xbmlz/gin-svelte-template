package core

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(conf Config, log Logger) Redis {
	addr := conf.Redis.Addr()

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       conf.Redis.DB,
		Password: conf.Redis.Password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Errorf("redis ping error: %s", err.Error())
	}

	log.Info("Redis connection established")

	return Redis{Client: client}
}
