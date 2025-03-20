package repository

import (
	"chat-app/common/database"

	"github.com/redis/go-redis/v9"
)

type IBaseRepository interface {
	SetDB()
	GetReadDB() *redis.Client
	GetWriteDB() *redis.Client
	BeginTxn() redis.Pipeliner
}

type BaseRepository struct {
	readDB  database.IRedisService
	writeDB database.IRedisService
}

func (b *BaseRepository) SetDB() {
	b.readDB = database.GetReadInstance()
	b.writeDB = database.GetWriteInstance()
}

func (b *BaseRepository) GetReadDB() *redis.Client {
	return b.readDB.GetConnection()
}

func (b *BaseRepository) GetWriteDB() *redis.Client {
	return b.writeDB.GetConnection()
}

func (b *BaseRepository) BeginTxn() redis.Pipeliner {
	return b.GetWriteDB().TxPipeline()
}
