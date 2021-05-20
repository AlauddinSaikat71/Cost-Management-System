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
	r.Use(CORSMiddleware())

	db.ConnectDatabase()
	db.InitModelDbContext()

	// `/uploads` should be same as UPLOAD_DIR value
	r.StaticFS("/uploads", http.Dir(filehandlers.UPLOAD_DIR))
	filehandlers.Seed()

	//routes for Users
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.GET("/users", controllers.FindUsers)
	r.DELETE("/users/:id", controllers.DeleteUser)

	//routes for Files
	r.POST("/uploads", filehandlers.SaveFileHandler)

	//routes for Attachment
	r.POST("/attachments", controllers.CreateAttachment)

	//routes for Cost
	r.POST("/costs", controllers.CreateCost)
	r.GET("/costs/:from/:to", controllers.FindCosts)
	r.PATCH("/costs/:id", controllers.UpdateCost)
	r.DELETE("/costs/:id", controllers.DeleteCost)

	//routes for CostAttachment
	r.POST("/costattachments", controllers.CreateCostAttachment)

	r.Run()
}
