package tasks

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type taskController struct {
	taskService TaskService
	ctx         *gin.Context
}

func NewTaskController(taskService TaskService, ctx *gin.Context) *taskController {
	return &taskController{taskService, ctx}
}

func (controller *taskController) Index(ctx *gin.Context) {
	data := controller.taskService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

func (controller *taskController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := controller.taskService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

func (controller *taskController) Create(ctx *gin.Context) {
	data, err := controller.taskService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    data,
		})
		ctx.Abort()
		return

	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}
