package usecase_test

import (
	"user_auth/domain"
	mock_domain "user_auth/domain/mock"
	"user_auth/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// SignUpの正常系テスト
func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリを利用
	mockur := mock_domain.NewMockIUserRepository(ctrl)

	// テストデータ
	user := domain.User{
		Email:    "test1@test.com",
		Password: "password1",
	}

	// CreateUserが呼び出されたときの挙動を設定
	mockur.EXPECT().CreateUser(gomock.Any()).DoAndReturn(func(u *domain.User) error {
		// パスワードを比較してハッシュ化されていることを確認
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("password1"))
		assert.NoError(t, err)
		u.ID = 1 // IDを設定
		return nil
	})
	
	uu := usecase.NewUserUsecase(mockur)
	result, err := uu.SignUp(user)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, user.Email, result.Email)
}

// LogInの正常系テスト
func TestLogIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// モックリポジトリを利用
	mockur := mock_domain.NewMockIUserRepository(ctrl)

	// テストデータ
	password := "password1"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	assert.NoError(t, err)
	user := domain.User{
		Email:    "test1@test.com",
		Password: string(hash), // ハッシュ化したパスワードを設定
	}

	// GetUserByEmailが呼び出されたときの挙動を設定
	mockur.EXPECT().GetUserByEmail(gomock.Any(), user.Email).DoAndReturn(
		func(u *domain.User, email string) error {
			*u = user // userを返す
			return nil
		})

	uu := usecase.NewUserUsecase(mockur)
	token, err := uu.LogIn(domain.User{
		Email:    user.Email,
		Password: password,
	})

	// 検証
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
