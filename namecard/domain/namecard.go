package domain

import (
	"time"
)

type Namecard struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Company   string    `json:"company" gorm:"not null"`
	Position  string    `json:"position" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserId    uint      `json:"user_id" gorm:"not null"`
}

// クライアントに返すデータの構造体
type NamecardResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Company   string    `json:"company" gorm:"not null"`
	Position  string    `json:"position" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}
