package model

import "time"

type User struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// クライアントに返すデータの構造体
type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
