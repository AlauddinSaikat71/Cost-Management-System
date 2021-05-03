package main

import (
	"costmanagement/controllers"
	"costmanagement/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()
	db.InitModelDbContext()

	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/users", controllers.FindUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
