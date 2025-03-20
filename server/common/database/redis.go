package database

import (
	"chat-app/common/constants/enums"
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisStore struct {
	client *redis.Client
}

var redisRead = RedisStore{}
var redisWrite = RedisStore{}

func InitRedisClient(ctx context.Context, dbType enums.AccessType) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADDR"), // "localhost:6379"
		Username: viper.GetString("REDIS_USER"), // "ADMIN"
		Password: viper.GetString("REDIS_PASS"), // "*****"
		DB:       0,                             // use default DB
	})

	try := 5
	for try > 0 {
		_, err := client.Ping(ctx).Result()
		if err == nil {
			break
		}

		try--
		if try == 0 {
			panic("redis connection failed" + err.Error())
		}
	}

	switch dbType {
	case enums.READ_ACCESS:
		redisRead.client = client
	case enums.WRITE_ACCESS:
		redisWrite.client = client
	}

	return &RedisStore{
		client: client,
	}
}

func GetReadInstance() *RedisStore {
	return &redisRead
}

func GetWriteInstance() *RedisStore {
	return &redisWrite
}

func (rs *RedisStore) GetConnection() *redis.Client {
	return rs.client
}

func (rs *RedisStore) Close() error {
	return rs.client.Close()
}
