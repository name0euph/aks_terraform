package repository

import (
	"context"
	"encoding/json"
	"go-rest-api/model"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

// userRepositoryのinterface
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	client *azcosmos.Client
	dbName string
	//db *gorm.DB
}

/*
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}\
}
// メールアドレスを元にユーザを取得
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// ユーザを作成
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
*/

func NewUserRepository(client *azcosmos.Client, dbName string) IUserRepository {
	return &userRepository{client, dbName}
}

// メールアドレスを元にユーザを取得
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	// DB名とコンテナ名を指定してコンテナを取得
	container, err := ur.client.NewContainer(ur.dbName, "Users")
	if err != nil {
		return err
	}

	pk := azcosmos.NewPartitionKeyString(email)
	itemResponse, err := container.ReadItem(context.Background(), pk, email, nil)
	if err != nil {
		return err
	}
	// JSONをuserオブジェクトにデシリアライズ
	err = json.Unmarshal(itemResponse.Value, user)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	// DB名とコンテナ名を指定してコンテナを取得
	container, err := ur.client.NewContainer(ur.dbName, "Users")
	if err != nil {
		return err
	}

	// userオブジェクトをJSONにエンコード
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	pk := azcosmos.NewPartitionKeyString(user.Email)
	_, err = container.CreateItem(context.Background(), pk, userData, nil)
	return err
}
