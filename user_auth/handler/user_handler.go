package handler

import (
	"user_auth/domain"
	uu "user_auth/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

// Handlerのインターフェース
type IUserHandler interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
}

type userHandler struct {
	uu uu.IUserUsecase
}

func NewUserHandler(uu uu.IUserUsecase) IUserHandler {
	return &userHandler{uu}
}

// SignUp
func (uh *userHandler) SignUp(c echo.Context) error {
	user := domain.User{} // 空のUser構造体を作成

	// リクエストボディをUser構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, err.Error()) // BINDに失敗した場合は400エラーを返す
	}

	// Signメソッドを呼び出してユーザ作成処理
	userRes, err := uh.uu.SignUp(user)
	if err != nil {
		return c.JSON(500, err.Error()) // サインアップ処理に失敗した場合は500エラーを返す
	}
	return c.JSON(201, userRes) // 成功した場合は201 Createdを返す
}

// LogIn
func (uh *userHandler) LogIn(c echo.Context) error {
	user := domain.User{} // 空のUser構造体を作成

	// リクエストボディをUser構造体にバインド
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, err.Error()) // BINDに失敗した場合は400エラーを返す
	}

	// LogInメソッドを呼び出してログイン処理
	token, err := uh.uu.LogIn(user)
	if err != nil {
		return c.JSON(500, err.Error()) // ログイン処理に失敗した場合は500エラーを返す
	}

	// トークンをクッキーにセット
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	// フロントエンドとバックエンドが異なるドメインの場合、SameSiteNoneModeを設定する
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(200) // 成功した場合は200 OKを返す
}

// LogOut
func (uh *userHandler) LogOut(c echo.Context) error {
	// トークンをクッキーから削除
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	// フロントエンドとバックエンドが異なるドメインの場合、SameSiteNoneModeを設定する
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(200) // 成功した場合は200 OKを返す
}