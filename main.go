package main

import (
	"github.com/gin-gonic/gin"
	"taskManagement.com/src/config"
	"taskManagement.com/src/router"

	t "taskManagement.com/src/modules/tasks"
	u "taskManagement.com/src/modules/users"
)

func main() {

	r := gin.Default()

	db := config.DB()

	db.AutoMigrate(&u.User{}, &t.Task{})

	router.Api(r, db)

	r.Run()

}
