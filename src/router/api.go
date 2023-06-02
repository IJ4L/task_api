package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	task "taskManagement.com/src/modules/tasks"
	user "taskManagement.com/src/modules/users"
)

var (
	ctx *gin.Context
)

func Api(router *gin.Engine, db *gorm.DB) {

	userReposiroty := user.NewUserRepository(db)
	userService := user.NewUserService(userReposiroty)
	userController := user.NewUserController(userService, ctx)

	taskReposiroty := task.NewTaskRepository(db)
	taskService := task.NewTaskService(taskReposiroty)
	taskController := task.NewTaskController(taskService, ctx)

	v1 := router.Group("/api")
	{
		v1.GET("/users", userController.Index)
		v1.GET("/users/:id", userController.Show)
		v1.POST("/users", userController.Create)

		v1.GET("/tasks", taskController.Index)
		v1.GET("/tasks/:id", taskController.Show)
		v1.POST("/tasks", taskController.Create)
	}

}
