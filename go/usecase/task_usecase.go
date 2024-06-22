package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

// taskUsecaseのinterfaceを定義
type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

// taskUsecaseの構造体を定義
type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

// taskUsecaseのコンストラクタ
func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	// taskUsecaseの構造体の実体を作成し、ポインタを返す
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	// 取得するタスク一覧のスライスを作成
	tasks := []model.Task{}
	// リクエストされたユーザIDを元にタスクを取得
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	// クライアントに返すデータの構造体のスライスを作成
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクを取得
	task := model.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	// クライアントに返すデータの構造体を作成
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// タスクを作成
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	// クライアントに返すデータの構造体を作成
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクを更新
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	// クライアントに返すデータの構造体を作成
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	// タスクを削除
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
