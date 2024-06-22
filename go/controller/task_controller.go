package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskContoller struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskContoller{tu}
}

func (tc *taskContoller) GetAllTasks(c echo.Context) error {
	// デコードされたJWTトークンからuserを取得し、クレームからuserIdを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskContoller) GetTaskById(c echo.Context) error {
	// デコードされたJWTトークンからuserを取得し、クレームからuserIdを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// パスパラメータからtaskIdを取得
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id) // string型からint型に変換

	// GetTaskByIdメソッドを呼び出してタスク取得処理
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskContoller) CreateTask(c echo.Context) error {
	// デコードされたJWTトークンからuserを取得し、クレームからuserIdを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// リクエストボディをTask構造体にバインド
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// taskのUserIdにuserから取得したuserIdをセット
	task.UserId = uint(userId.(float64))
	// CreateTaskメソッドを呼び出してタスク作成処理
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskContoller) UpdateTask(c echo.Context) error {
	// デコードされたJWTトークンからuserを取得し、クレームからuserIdを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// パスパラメータからtaskIdを取得
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id) // string型からint型に変換

	// リクエストボディをTask構造体にバインド
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// UpdateTaskメソッドを呼び出してタスク更新処理
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskContoller) DeleteTask(c echo.Context) error {
	// デコードされたJWTトークンからuserを取得し、クレームからuserIdを取得
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// パスパラメータからtaskIdを取得
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id) // string型からint型に変換

	// DeleteTaskメソッドを呼び出してタスク削除処理
	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
