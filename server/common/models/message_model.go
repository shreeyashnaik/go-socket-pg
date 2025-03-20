package models

import (
	"chat-app/common/constants/enums"
	"time"
)

type Message struct {
	Username  string            `json:"username"`
	ChatID    string            `json:"chat_id"`
	Content   string            `json:"content"`
	Timestamp time.Time         `json:"timestamp"`
	Type      enums.MessageType `json:"type"`
}
