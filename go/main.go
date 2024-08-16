package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

// @title        Todo API
// @version      1.0
// @description  This is a simple todo API.

// @contact.name   name0euph
// @contact.email  name0euph@gmail.com

// @host      localhost:8080
// @BasePath  /
func main() {
	db := db.NewDB()
	// 依存性の注入
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
