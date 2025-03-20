package errors

import "errors"

var (
	ErrChatRoomNotFound      = errors.New("chat room not found")
	ErrMaxUsersReached       = errors.New("maximum number of users reached in the chat room")
	ErrUsernameAlreadyExists = errors.New("username already exists in the chat room")
)
