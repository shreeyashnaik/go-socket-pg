package service

import (
	"chat-app/common/constants"
	domainerrors "chat-app/common/errors"
	"chat-app/common/models"
	"chat-app/common/schemas"
	"chat-app/common/utils"
	"chat-app/httpserver/repository"
	"context"
)

type IChatService interface {
	CreateChatRoom(ctx context.Context, chatRoomRequest schemas.CreateChatRequest) (string, error)
	AddUserToChatRoom(ctx context.Context, chatRoomID string, username string) error
}

type ChatService struct {
	chatRepo repository.IChatRepository
	userRepo repository.IUserRepository
}

func InitChatService() *ChatService {
	return &ChatService{
		chatRepo: repository.InitChatRepository(),
		userRepo: repository.InitUserRepository(),
	}
}

func (svc *ChatService) CreateChatRoom(ctx context.Context, chatRoomRequest schemas.CreateChatRequest) (string, error) {

	// Prepare the Chat Room Model
	chatRoomModel := models.ChatRoom{
		ID:       utils.RandString(constants.CHAT_ROOM_ID_LENGTH),
		Name:     chatRoomRequest.ChatRoomName,
		Admin:    chatRoomRequest.AdminName,
		IsActive: true,
	}

	// Start a new transaction
	txn := svc.chatRepo.BeginTxn()

	// Task 1: Create Chat Room
	if err := svc.chatRepo.CreateChatRoom(ctx, txn, getChatKey(chatRoomModel.ID), chatRoomModel); err != nil {
		txn.Discard()
		return "", err
	}

	// Task 2: Add Admin to list of Users
	if err := svc.userRepo.CreateUser(ctx, txn, getUsersKey(chatRoomModel.ID), chatRoomModel.Admin); err != nil {
		txn.Discard()
		return "", err
	}

	// Commit the transaction
	if _, err := txn.Exec(ctx); err != nil {
		return "", err
	}

	return chatRoomModel.ID, nil
}

func (svc *ChatService) AddUserToChatRoom(ctx context.Context, chatRoomID string, username string) error {

	// Fetch all existing users in the chat room
	usersByRoom, err := svc.userRepo.GetUsers(ctx, nil, getUsersKey(chatRoomID))
	if err != nil {
		return err
	}

	if len(usersByRoom) == 0 {
		return domainerrors.ErrChatRoomNotFound
	}

	// Check if the chat room has reached the maximum allowed users
	if len(usersByRoom) >= constants.MAX_ALLOWED_CHAT_ROOM_USERS {
		return domainerrors.ErrMaxUsersReached
	}

	// Check if the username already exists in the chat room
	if utils.Contains(usersByRoom, username) {
		return domainerrors.ErrUsernameAlreadyExists
	}

	// Add the user to the chat room
	if err := svc.userRepo.CreateUser(ctx, nil, getUsersKey(chatRoomID), username); err != nil {
		return err
	}

	return nil
}

func getChatKey(chatRoomID string) string {
	return constants.ROOM_PREFIX + "_" + chatRoomID + "_" + "details"
}

func getUsersKey(chatRoomID string) string {
	return constants.ROOM_PREFIX + "_" + chatRoomID + "_" + "users"
}
