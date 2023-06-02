package users

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() []User
	GetOne(id int) User
	Create(user User) (*User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (repo *UserRepositoryImpl) GetAll() []User {
	var users []User

	_ = repo.db.Model(&users).Preload("Task").Find(&users).Error

	return users
}

func (repo *UserRepositoryImpl) GetOne(id int) User {
	var user User

	_ = repo.db.Model(&user).Preload("Task").Where("id = ?", id).First(&user)

	return user
}

func (repo *UserRepositoryImpl) Create(user User) (*User, error) {
	result := repo.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
