package handler_test

import (
	"user_auth/domain"
	"user_auth/handler"
	mock_usecase "user_auth/usecase/mock"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// SignUpの正常系テスト
func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockuu := mock_usecase.NewMockIUserUsecase(ctrl)

	// テストデータ
	user := domain.User{
		Email:    "test1@test.com",
		Password: "password1",
	}
	userRes := domain.ResUser{
		ID:    0,
		Email: user.Email,
	}

	// SignUpが呼び出されたときの挙動を設定
	mockuu.EXPECT().SignUp(user).Return(userRes, nil)

	// テスト対象のハンドラを作成
	e := echo.New()
	reqBody, err := json.Marshal(user)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uh := handler.NewUserHandler(mockuu)
	err = uh.SignUp(c)

	// 検証
	assert.NoError(t, err)         // エラーが発生していないこと
	assert.Equal(t, http.StatusCreated, rec.Code) // ステータスコード201が返却されていること

	expectedResBody, err := json.Marshal(userRes)
	assert.NoError(t, err)
	reqBody, err = io.ReadAll(rec.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, string(expectedResBody), string(reqBody)) // レスポンスのボディが期待通りであること
}

// LogInの正常系テスト
func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockuu := mock_usecase.NewMockIUserUsecase(ctrl)

	// テストデータ
	user := domain.User{
		Email: "test1@test.com",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	assert.NoError(t, err)

	// LogInが呼び出されたときの挙動を設定
	mockuu.EXPECT().LogIn(user).Return(tokenString, nil)

	// テスト対象のハンドラを作成
	e := echo.New()
	reqBody, err := json.Marshal(user)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(reqBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uh := handler.NewUserHandler(mockuu)
	err = uh.LogIn(c)

	// 検証
	assert.NoError(t, err)                   // エラーが発生していないこと
	assert.Equal(t, http.StatusOK, rec.Code) // ステータスコード200が返却されていること

	cookie := rec.Result().Cookies()                    // レスポンスヘッダからクッキーを取得
	assert.NotNil(t, cookie)                            // クッキーがセットされていること
	assert.Equal(t, "token", cookie[0].Name)            // クッキー名が"token"であること
	assert.Equal(t, tokenString, cookie[0].Value)       // クッキーの値がトークン文字列と一致すること
	assert.True(t, cookie[0].Expires.After(time.Now())) // クッキーの有効期限が未来であること
}

// LogOutの正常系テスト
func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockuu := mock_usecase.NewMockIUserUsecase(ctrl)

	// テスト対象のハンドラを作成
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	uh := handler.NewUserHandler(mockuu)
	err := uh.LogOut(c)

	// 検証
	assert.NoError(t, err)                   // エラーが発生していないこと
	assert.Equal(t, http.StatusOK, rec.Code) // ステータスコード200が返却されていること

	cookie := rec.Result().Cookies()                     // レスポンスヘッダからクッキーを取得
	assert.NotNil(t, cookie)                             // クッキーがセットされていること
	assert.Equal(t, "token", cookie[0].Name)             // クッキー名が"token"であること
	assert.Equal(t, "", cookie[0].Value)                 // クッキーの値が空であること
	assert.True(t, cookie[0].Expires.Before(time.Now())) // クッキーの有効期限が過去であること
}
