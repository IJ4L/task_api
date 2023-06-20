package tasks

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "taskManagement.com/src/modules/tasks/dto"
)

type TaskService interface {
	GetAll() []Task
	GetById(id int) Task
	Create(ctx *gin.Context) (*Task, error)
	Update(ctx *gin.Context) (*Task, error)
	Delete(ctx *gin.Context) (*Task, error)
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
		Status: false,
	}

	result, err := service.taskRepository.Create(task)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (service *TaskServiceImpl) Update(ctx *gin.Context) (*Task, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateTaskInputDto

	if err := ctx.ShouldBind(&input); err != nil {
		return nil, err
	}

	viladate := validator.New()

	err := viladate.Struct(&input)

	if err != nil {
		return nil, err
	}

	task := Task{
		ID:   int64(id),
		Task: input.Task,
		Time: input.Time,
	}

	result, err := service.taskRepository.Update(task)

	if result != nil {
		return nil, err
	}

	return result, nil
}

func (service *TaskServiceImpl) Delete(ctx *gin.Context) (*Task, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	task := Task{
		ID: int64(id),
	}

	result, err := service.taskRepository.Delete(task)

	if result != nil {
		return nil, err
	}

	return result, nil
}
