package models

type ChatRoom struct {
	ID       string `json:"id" redis:"-"`
	Name     string `json:"name" redis:"name"`
	Admin    string `json:"admin" redis:"admin"`
	IsActive bool   `json:"is_active" redis:"is_active"`
}
