package controller

import (
	domainerrors "chat-app/common/errors"
	"chat-app/common/schemas"
	"chat-app/httpserver/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ChatHandler struct {
	BaseController
	service service.IChatService
}

func InitChatHandler() *ChatHandler {
	return &ChatHandler{
		service: service.InitChatService(),
	}
}

func (c *ChatHandler) CreateChatRoom(w http.ResponseWriter, r *http.Request) {

	// Validate the request Method
	if r.Method != http.MethodPost {
		c.ResponseMethodNotAllowed(w, r.Method+" method not allowed", nil)
		return
	}

	// Decode the request payload
	payload := schemas.CreateChatRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.ResponseBadRequest(w, "failed to decode payload", nil)
		return
	}

	// Validate the request payload
	if err := validator.New().Struct(payload); err != nil {
		c.ResponseBadRequest(w, "invalid params", nil)
		return
	}

	// Service call
	chatRoomID, err := c.service.CreateChatRoom(r.Context(), payload)
	if err != nil {
		if err == domainerrors.ErrChatRoomNotFound || err == domainerrors.ErrUsernameAlreadyExists || err == domainerrors.ErrMaxUsersReached {
			c.ResponseBadRequest(w, err.Error(), nil)
			return
		}
		c.ResponseInternalServerError(w, "server error", nil)
		return
	}

	// Return success response
	c.ResponseCreated(w, map[string]any{
		"chat_room_id": chatRoomID,
	})
}

func (c *ChatHandler) AddUser(w http.ResponseWriter, r *http.Request) {

	// Validate the request Method
	if r.Method != http.MethodPost {
		c.ResponseMethodNotAllowed(w, r.Method+" method not allowed", nil)
		return
	}

	// Decode the request payload
	payload := schemas.AddUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.ResponseBadRequest(w, "failed to decode payload", nil)
		return
	}

	// Extract the chatID from the URL
	payload.ChatRoomID = chi.URLParam(r, "chatID")

	// Validate the request payload
	if err := validator.New().Struct(payload); err != nil {
		c.ResponseBadRequest(w, "invalid params", nil)
		return
	}

	// Service call
	if err := c.service.AddUserToChatRoom(r.Context(), payload.ChatRoomID, payload.Username); err != nil {
		if err == domainerrors.ErrChatRoomNotFound || err == domainerrors.ErrUsernameAlreadyExists || err == domainerrors.ErrMaxUsersReached {
			c.ResponseBadRequest(w, err.Error(), nil)
			return
		}
		c.ResponseInternalServerError(w, "server error", nil)
		return
	}

	// Return success response
	c.ResponseCreated(w, nil)
}
