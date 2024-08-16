package main

import (
	"user_auth/domain"
	"user_auth/handler"
	"user_auth/infrastructure"
	"user_auth/usecase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// ユーザ作成のテスト
func TestCreateUser(t *testing.T) {
	// テスト用のデータ
	user := domain.User{
		Email:    "test@test.com",
		Password: "password",
	}

	// テスト用のデータをJSONに変換
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// リクエストを作成
	req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(string(userJSON))) // Request
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)                            // Content-Type: application/json

	rec := httptest.NewRecorder() // ResponseWriter

	db := infrastructure.NewDB() // DB接続
	ur := domain.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)
	e := NewRouter(uh)
	e.ServeHTTP(rec, req) // Handle

	assert.Equal(t, http.StatusCreated, rec.Code) // 期待ステータス: 201 Created
}

// ユーザログインのテスト
func TestLoginUser(t *testing.T) {
	// テスト用のデータ
	user := domain.User{
		Email:    "test@test.com",
		Password: "password",
	}

	// テスト用のデータをJSONに変換
	userJSON, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// リクエストを作成
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(userJSON))) // Request
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)                           // Content-Type: application/json

	rec := httptest.NewRecorder() // ResponseWriter

	db := infrastructure.NewDB() // DB接続
	ur := domain.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)
	e := NewRouter(uh)
	e.ServeHTTP(rec, req) // Handle

	assert.Equal(t, http.StatusOK, rec.Code) // 期待ステータス: 200 OK
}

// ユーザログアウトのテスト
func TestLogoutUser(t *testing.T) {
	// リクエストを作成
	req := httptest.NewRequest("POST", "/logout", nil) // Request

	rec := httptest.NewRecorder() // ResponseWriter

	db := infrastructure.NewDB() // DB接続
	ur := domain.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)
	e := NewRouter(uh)
	e.ServeHTTP(rec, req) // Handle

	assert.Equal(t, http.StatusOK, rec.Code) // 期待ステータス: 200 OK
}
