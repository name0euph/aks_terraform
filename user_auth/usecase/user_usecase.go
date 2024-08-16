package usecase

import (
	"user_auth/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user domain.User) (domain.ResUser, error)
	LogIn(user domain.User) (string, error)
}

type userUsecase struct {
	ur domain.IUserRepository
}

// 依存性の注入
func NewUserUsecase(ur domain.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

// サインアップ
func (uu *userUsecase) SignUp(user domain.User) (domain.ResUser, error) {
	// パスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return domain.ResUser{}, err
	}

	// DBに保存するユーザ情報を作成
	newUser := domain.User{
		Email:    user.Email,
		Password: string(hash),
	}

	if err := uu.ur.CreateUser(&newUser); err != nil {
		return domain.ResUser{}, err
	}

	resUser := domain.ResUser{
		ID:   newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

// ログイン
func (uu *userUsecase) LogIn(user domain.User) (string, error) {
	loginUser := domain.User{}

	// DBからEmailを元にユーザを取得
	if err := uu.ur.GetUserByEmail(&loginUser, user.Email); err != nil {
		return "", err
	}

	// パスワードの照合
	if err := bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	// JWTトークンの定義
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": loginUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	// トークン文字列を返す
	return tokenString, nil
}
