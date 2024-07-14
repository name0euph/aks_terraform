package main

import (
	"bytes"
	"context"
	"encoding/json"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
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

	Data := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "user3@test.com",
		Password: "Password",
	}

	jsonData, err := json.Marshal(Data)
	if err != nil {
		t.Fatalf("JSONのエンコードに失敗しました: %s", err)
	}

	// サインアップ
	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		e.URL+"/signup",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		t.Fatalf("リクエストの送信に失敗しましたs: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, _ := client.Do(req)
	if err != nil {
		t.Fatalf("リクエストの送信に失敗しましたs: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("ステータスコード: %v", res.Status)
	}
}
