package domain

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

type ResUser struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}