package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	task := []entity.Task{}

	tasks := r.db.WithContext(ctx).Where("user_id = ?", id).Find(&task)
	if tasks.Error != nil {
		return []entity.Task{}, tasks.Error
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	tasks := entity.Task{}
	tsk := r.db.WithContext(ctx).Create(&task)

	if tsk.Error != nil {
		return 0, tsk.Error
	}
	tsk = r.db.WithContext(ctx).Where("category_id = ?", task.CategoryID).Order("id DESC").First(&tasks)
	if tsk.Error != nil {
		return 0, tsk.Error
	}
	return tasks.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	task := entity.Task{}
	tasks := r.db.WithContext(ctx).Where("id = ?", id).First(&task)
	if tasks.Error != nil {
		return entity.Task{}, tasks.Error
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	task := []entity.Task{}
	tasks := r.db.WithContext(ctx).Where("category_id = ?", catId).Find(&task)
	if tasks.Error != nil {
		return []entity.Task{}, tasks.Error
	}
	return task, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	tasks := r.db.WithContext(ctx).Model(&entity.Task{}).Where("id = ?", task.ID).Updates(task)
	if tasks.Error != nil {
		return tasks.Error
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	task := r.db.WithContext(ctx).Delete(&entity.Task{}, id)
	if task.Error != nil {
		return task.Error
	}
	return nil // TODO: replace this
}
