package schemas

type CreateChatRequest struct {
	ChatRoomName string `json:"chat_room_name" validate:"required"`
	AdminName    string `json:"admin_name" validate:"required"`
}

type AddUserRequest struct {
	ChatRoomID string `json:"chat_room_id" validate:"required"`
	Username   string `json:"username" validate:"required"`
}

type GenericHTTPResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}
