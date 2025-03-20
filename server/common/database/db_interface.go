package database

import "github.com/redis/go-redis/v9"

type IRedisService interface {
	Close() error
	GetConnection() *redis.Client
}

// type ISqlService interface {
// 	Close() error
// 	GetConnection() *gorm.DB
// }
