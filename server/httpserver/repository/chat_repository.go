package repository

import (
	"chat-app/common/models"
	"context"

	"github.com/redis/go-redis/v9"
)

type IChatRepository interface {
	IBaseRepository

	CreateChatRoom(ctx context.Context, txn redis.Cmdable, key string, chatRoomModel models.ChatRoom) error
}

type ChatRepository struct {
	BaseRepository
}

func InitChatRepository() *ChatRepository {
	repo := &ChatRepository{}
	repo.SetDB()
	return repo
}

func (repo *ChatRepository) CreateChatRoom(ctx context.Context, txn redis.Cmdable, key string, chatRoomModel models.ChatRoom) error {
	if txn == nil {
		txn = repo.GetWriteDB()
	}

	_, err := txn.HSet(ctx, key, chatRoomModel).Result()
	return err
}
