package main

import (
	"costmanagement/controllers"
	"costmanagement/db"
	"costmanagement/filehandlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	CORSMiddleware()

	db.ConnectDatabase()
	db.InitModelDbContext()

	// `/uploads` should be same as UPLOAD_DIR value
	r.StaticFS("/uploads", http.Dir(filehandlers.UPLOAD_DIR))
	filehandlers.Seed()

	//routes for users
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/users", controllers.FindUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)

	//routes for files
	r.POST("/uploads", filehandlers.SaveFileHandler)

	r.Run()
}
