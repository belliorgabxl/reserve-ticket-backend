package redisx

import (
	"context"
	"strconv"

	"github.com/belliorgabxl/reserve-ticket-backend/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg config.Config) (*redis.Client, error) {

	db, _ := strconv.Atoi(cfg.RedisDB)

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       db,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
