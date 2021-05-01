package main

import (
	"costmanagement/controllers"
	"costmanagement/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)

	r.Run()
}
