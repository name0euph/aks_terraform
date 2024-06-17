package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// 外部キー制約、Userが削除されたらTaskも削除される
	User   User `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId uint `json:"user_id" gorm:"not null"`
}

// クライアントに返すデータの構造体
type TaskResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
