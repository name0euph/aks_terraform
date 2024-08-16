package main

import (
	"user_auth/domain"
	"user_auth/handler"
	"user_auth/infrastructure"
	"user_auth/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(uh handler.IUserHandler) *echo.Echo {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Health Check: OK")
	})
	e.POST("/signup", uh.SignUp)
	e.POST("/login", uh.LogIn)
	e.POST("/logout", uh.LogOut)

	return e
}

func main() {
	db := infrastructure.NewDB()
	ur := domain.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uu)
	e := NewRouter(uh)

	e.Logger.Fatal(e.Start(":8080"))
}
