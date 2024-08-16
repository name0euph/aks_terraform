package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
	HealthCheck(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// 依存性の注入
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// SignUp godoc
// @Summary      サインアップ
// @Description  新しいユーザーを作成する
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        username  body      string  true  "ユーザー名"
// @Param        password  body      string  true  "パスワード"
// @Success      200       {object}  map[string]interface{}
// @Failure      400       {object}  map[string]interface{}
// @Failure      500       {object}  map[string]interface{}
// @Router       /signup [post]
func (uc *userController) SignUp(c echo.Context) error {
	// リクエストボディをUser構造体にバインド
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Signメソッドを呼び出してユーザ作成処理
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, userRes)
}

// LogIn godoc
// @Summary      ログイン
// @Description  既存のユーザーでログインする
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        username  body      string  true  "ユーザー名"
// @Param        password  body      string  true  "パスワード"
// @Success      200       {object}  map[string]interface{}
// @Failure      400       {object}  map[string]interface{}
// @Failure      500       {object}  map[string]interface{}
// @Router       /login [post]
func (uc *userController) LogIn(c echo.Context) error {
	// リクエストボディをUser構造体にバインド
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Loginメソッドを呼び出してトークンを取得
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// トークンをクッキーにセット
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	// フロントエンドとバックエンドが異なるドメインの場合、SameSiteNoneModeを設定する
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// LogOut godoc
// @Summary      ログアウト
// @Description  ユーザがログアウトする
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        username  body      string  true  "ユーザー名"
// @Param        password  body      string  true  "パスワード"
// @Success      200       {object}  map[string]interface{}
// @Router       /logout [post]
func (uc *userController) LogOut(c echo.Context) error {
	// 値を空かつ有効期限をtime.Now()に設定してクッキーを削除
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

// CSRF godoc
// @Summary      CSRF
// @Description  CSRFトークンを取得する
// @Tags         others
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /csrf [get]
func (uc *userController) CsrfToken(c echo.Context) error {
	// CSRFトークンを取得
	token := c.Get("csrf").(string)

	// csrf_tokenにトークンをセットしてクライアントに返却
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}

// HealthCheck godoc
// @Summary      正常性エンドポイント
// @Description  正常性を確認するためのエンドポイント
// @Tags         others
// @Accept       json
// @Produce      json
// @Success      200  {string}  string  "OK"
// @Router       /health [get]
func (uc *userController) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
