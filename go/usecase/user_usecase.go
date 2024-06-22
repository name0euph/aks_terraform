package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// usecaseのinterface
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	// repositoryのinterfaceに依存する
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// 依存性の注入
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

// サインアップ処理を定義する関数
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	// バリデーションチェック
	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}

	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// ハッシュ化に失敗した場合、0値とエラーを返す
	if err != nil {
		return model.UserResponse{}, err
	}

	// DBに保存するユーザ情報を作成
	newUser := model.User{
		Email:    user.Email,
		Password: string(hash),
	}

	// ユーザ作成処理、エラーがあれば0値とエラーを返す
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	// レスポンス用のUser構造体を作成
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	// バリデーションチェック
	if err := uu.uv.UserValidate(user); err != nil {
		return "", err
	}

	storedUser := model.User{}

	// 入力されたEmailがDBに存在するかを検証
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	// パスワードが一致するかを検証
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// JWTトークンの定義
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		// トークンの有効期限は12時間
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})
	// tokenの生成
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
