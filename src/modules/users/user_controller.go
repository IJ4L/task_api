package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService UserService
	ctx         *gin.Context
}

func NewUserController(userService UserService, ctx *gin.Context) *UserController {
	return &UserController{userService, ctx}
}

func (controller *UserController) Index(ctx *gin.Context) {
	data := controller.userService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

func (controller *UserController) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data := controller.userService.GetById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

func (controller *UserController) Create(ctx *gin.Context) {
	data, err := controller.userService.Create(ctx)

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
