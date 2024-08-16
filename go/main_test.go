package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthCheck(t *testing.T) {

	// ルータの設定
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := httptest.NewServer(router.NewRouter(userController, taskController))
	defer e.Close()

	r, err := http.Get(e.URL + "/health")
	if err != nil {
		t.Fatalf("リクエストの送信に失敗しました: %s", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("ステータスコード: %d", r.StatusCode)
	}
}

func TestUserHandler(t *testing.T) {

	// ルータの設定
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := httptest.NewServer(router.NewRouter(userController, taskController))
	defer e.Close()

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// CSRFトークンの取得
	csrfToken, err := getCsrfToken(&client, e.URL)
	fmt.Printf("%s\n", csrfToken)
	if err != nil {
		t.Fatalf("CSRFトークンの取得に失敗しました: %s", err)
	}

	jsonData := map[string]string{
		"email":    "user3@test.com",
		"password": "Password",
	}
	reqBody, err := json.Marshal(jsonData)
	if err != nil {
		t.Fatalf("%s", err)
	}

	// サインアップ
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		e.URL+"/signup",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		t.Fatalf("リクエストの作成に失敗しましたs: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-TOKEN", csrfToken) // CSRFトークンをヘッダにセット

	fmt.Printf("リクエストヘッダ: %v\n", req.Header)

	res, err := client.Do(req)
	if err != nil {
		t.Fatalf("リクエストの送信に失敗しましたs: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("ステータスコード: %v", res.Status)
		body, _ := io.ReadAll(res.Body)
		t.Errorf("レスポンスボディ: %s", string(body))
	}
}

func getCsrfToken(client *http.Client, url string) (string, error) {
	req, err := http.NewRequest("GET", url+"/csrf", nil)
	if err != nil {
		return "", fmt.Errorf("CSRFトークン取得リクエストの作成に失敗しました: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("CSRFトークン取得リクエストの送信に失敗しました: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("CSRFトークンの取得に失敗しました。ステータスコード: %d", res.StatusCode)
	}
	defer res.Body.Close()

	var responseBody map[string]string
	if err := json.NewDecoder(res.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("レスポンスボディのデコードに失敗しました: %w", err)
	}

	csrfToken, ok := responseBody["csrf_token"]
	if !ok {
		return "", fmt.Errorf("CSRFトークンがレスポンスに含まれていません。")
	}

	return csrfToken, nil
}
