package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type IUserRepository interface {
	IBaseRepository

	CreateUser(ctx context.Context, txn redis.Cmdable, usersKey string, username string) error
	DeleteUser(ctx context.Context, txn redis.Cmdable, usersKey string, username string) error
	GetUsers(ctx context.Context, txn redis.Cmdable, usersKey string) ([]string, error)
}

type UserRepository struct {
	BaseRepository
}

func InitUserRepository() *UserRepository {
	repo := &UserRepository{}
	repo.SetDB()
	return repo
}

func (repo *UserRepository) CreateUser(ctx context.Context, txn redis.Cmdable, usersKey string, username string) error {
	if txn == nil {
		txn = repo.GetWriteDB()
	}

	_, err := txn.SAdd(ctx, usersKey, username).Result()
	return err
}

func (repo *UserRepository) DeleteUser(ctx context.Context, txn redis.Cmdable, usersKey string, username string) error {
	if txn == nil {
		txn = repo.GetWriteDB()
	}

	_, err := txn.SRem(ctx, usersKey, username).Result()
	return err
}

func (repo *UserRepository) GetUsers(ctx context.Context, txn redis.Cmdable, usersKey string) ([]string, error) {
	if txn == nil {
		txn = repo.GetWriteDB()
	}

	return txn.SMembers(ctx, usersKey).Result()
}
