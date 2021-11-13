package database

import (
	"auth-control/configurations"
	appErrors "auth-control/errors"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Redis *redis.Client
}

func CreateDatabase() (IDatabase, error) {
	redisHost := fmt.Sprintf("%s:%s", configurations.Envs.RedisHost, configurations.Envs.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: configurations.Envs.RedisPassword,
		DB:       configurations.Envs.RedisDatabase,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		return nil, err
	}

	return &RedisClient{Redis: client}, nil
}

func (r *RedisClient) Set(ctx context.Context, key string, value string, ex time.Duration) (string, appErrors.ErrorResponse) {
	var appError appErrors.ErrorResponse
	_, err := r.Redis.Set(ctx, key, value, ex).Result()
	if err != nil {
		appError = appErrors.DatabaseOperationError("Error to set info in database")
		return "", appError
	}
	return key, appError
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, appErrors.ErrorResponse) {
	var appError appErrors.ErrorResponse
	value, err := r.Redis.Get(ctx, key).Result()
	if err != nil {
		log.Println(fmt.Printf("Error to get token - %s", err.Error()))
		appError = appErrors.DatabaseOperationError("Error to get info from database")
		return "", appError
	}
	return value, appError
}

func (r *RedisClient) Delete(ctx context.Context, key string) appErrors.ErrorResponse {
	var appError appErrors.ErrorResponse
	_, err := r.Redis.Del(ctx, key).Result()
	if err != nil {
		appError = appErrors.DatabaseOperationError("Error to delete info from database")
		return appError
	}
	return appError
}
