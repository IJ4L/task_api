package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "taskManagement.com/src/modules/users/dto"
)

type UserService interface {
	GetAll() []User
	GetById(id int) User
	Create(ctx *gin.Context) (*User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &UserServiceImpl{userRepository}
}

func (service *UserServiceImpl) GetAll() []User {
	return service.userRepository.GetAll()
}

func (service *UserServiceImpl) GetById(id int) User {
	return service.userRepository.GetOne(id)
}

func (service *UserServiceImpl) Create(ctx *gin.Context) (*User, error) {
	var input dto.CreateUserInputDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	user := User{
		Name:  input.Name,
		Email: input.Email,
	}

	result, err := service.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}
