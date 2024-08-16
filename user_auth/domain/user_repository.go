package domain

import (
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(user *User, email string) error
	DeleteUser(user *User) error
}

type UserRepository struct {
	db *gorm.DB
}

// データベースの接続を外側から受け取る
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

// ユーザーを作成する
func (ur *UserRepository) CreateUser(user *User) error {
	// ユーザーを作成
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// メールアドレスを元にユーザを取得する
func (ur *UserRepository) GetUserByEmail(user *User, email string) error {
	// DBからメールアドレスを元にユーザを取得
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザーを削除する
func (ur *UserRepository) DeleteUser(user *User) error {
	// ユーザーを削除
	if err := ur.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
