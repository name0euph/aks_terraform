package router

import (
	"go-rest-api/controller"
	"os"

	_ "go-rest-api/docs"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")}, // フロントエンドのURLを許可
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken}, // 許可するHTTPヘッダー
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // 許可するHTTPメソッド
		AllowCredentials: true,                                     // cookieを許可
	}))
	/*
		e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
			CookiePath:     "/",
			CookieDomain:   os.Getenv("API_DOMAIN"),
			CookieHTTPOnly: true,
			//CookieSameSite: http.SameSiteNoneMode,
			CookieSameSite: http.SameSiteDefaultMode,
			//CookieMaxAge: 60,
		}))
	*/

	// ルーティングの設定
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	//e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token", // cookieに保存されたトークンを取得
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)

	// 正常性エンドポイント
	e.GET("/health", uc.HealthCheck)

	// OpenAPIドキュメントエンドポイント
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return e
}
