package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "taskManagement.com/src/modules/tasks/dto"
)

type TaskService interface {
	GetAll() []Task
	GetById(id int) Task
	Create(ctx *gin.Context) (*Task, error)
}

type TaskServiceImpl struct {
	taskRepository TaskRepository
}

func NewTaskService(taskRepository TaskRepository) TaskService {
	return &TaskServiceImpl{taskRepository}
}

func (service *TaskServiceImpl) GetAll() []Task {
	return service.taskRepository.GetAll()
}

func (service *TaskServiceImpl) GetById(id int) Task {
	return service.taskRepository.GetOne(id)
}

func (service *TaskServiceImpl) Create(ctx *gin.Context) (*Task, error) {
	var input dto.CreateTaskInputDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	task := Task{
		Task:   input.Task,
		Time:   input.Time,
		UserID: input.UserID,
	}

	result, err := service.taskRepository.Create(task)

	if err != nil {
		return nil, err
	}

	return result, nil

}
