package main

import (
	"costmanagement/models"
 
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Run()
}
