package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// userRepositoryのinterface
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// メールアドレスを元にユーザを取得
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// DBからメールアドレスを元にユーザを取得
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// ユーザをDBに作成
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
