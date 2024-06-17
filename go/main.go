package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"os"
)

func main() {
	db := db.NewDB()
	// 依存性の注入
	userRepository := repository.NewUserRepository(db, os.Getenv("COSMOS_DB_NAME"))
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
