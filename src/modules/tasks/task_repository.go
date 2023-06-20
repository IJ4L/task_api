package tasks

import "gorm.io/gorm"

type TaskRepository interface {
	GetAll() []Task
	GetOne(id int) Task
	Create(task Task) (*Task, error)
	Update(task Task) (*Task, error)
	Delete(task Task) (*Task, error)
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

	_ = repo.db.Find(&task, id).Error

	return task
}

func (repo *TaskRepositoryImpl) Create(task Task) (*Task, error) {
	result := repo.db.Create(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (repo *TaskRepositoryImpl) Update(task Task) (*Task, error) {
	result := repo.db.Model(&Task{}).Where("id = ?", task.ID).Updates(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}

func (repo *TaskRepositoryImpl) Delete(task Task) (*Task, error) {
	result := repo.db.Delete(&task)

	if result.Error != nil {
		return nil, result.Error
	}

	return &task, nil
}
