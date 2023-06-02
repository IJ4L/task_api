package tasks

import "gorm.io/gorm"

type TaskRepository interface {
	GetAll() []Task
	GetOne(id int) Task
	Create(task Task) (*Task, error)
}

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &TaskRepositoryImpl{db}
}

func (repo *TaskRepositoryImpl) GetAll() []Task {
	var tasks []Task

	_ = repo.db.Model(&tasks).Find(&tasks).Error

	return tasks
}

func (repo *TaskRepositoryImpl) GetOne(id int) Task {
	var task Task

	_ = repo.db.Model(&task).Preload("User").First(&task).Error

	return task
}

func (repo *TaskRepositoryImpl) Create(task Task) (*Task, error) {
	result := repo.db.Create(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}
