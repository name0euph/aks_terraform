package main

import (
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

	// タイムアウトを設定したHTTPクライアントを使用
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	r, err := client.Get(e.URL + "/health")
	if err != nil {
		t.Fatalf("リクエストの送信に失敗しました: %s", err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", r.StatusCode)
	}
}
