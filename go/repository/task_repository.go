package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// taskRepositoryのinterfaceを定義
type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId uint) error
	GetTaskById(tasks *model.Task, userId uint, taskId uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

// taskRepositoryの構造体を定義
type taskRepository struct {
	db *gorm.DB
}

// taskRepositoryのコンストラクタを定義
func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(task *[]model.Task, userId uint) error {
	// UserテーブルとTaskテーブルを結合して、ユーザIDを元にタスクを取得し、CreatedAtでソート
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	// UserテーブルとTaskテーブルを結合して、ユーザIDとタスクIDを元にタスクを取得
	if err := tr.db.Joins("User").Where("user_id=?, userId").First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	// タスクを作成
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	// タスクIDとユーザIDを元にタスクを取得し、更新
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	// 何かしらのエラーが発生した場合、エラーを返す
	if result.Error != nil {
		return result.Error
	}
	// 更新対象のレコードが存在しない場合、object dose not existを返す
	if result.RowsAffected < 1 {
		return fmt.Errorf(("object dose not exist"))
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	// タスクIDとユーザIDを元にタスクを取得し、削除
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	// 何かしらのエラーが発生した場合、エラーを返す
	if result.Error != nil {
		return result.Error
	}
	// 削除対象のレコードが存在しない場合、object dose not existを返す
	if result.RowsAffected < 1 {
		return fmt.Errorf("object dose not exist")
	}
	return nil
}
