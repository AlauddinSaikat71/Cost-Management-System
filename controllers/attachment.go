package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /attachment
// create new attachment
func CreateAttachment(c *gin.Context) {
	//validate input
	var input dtos.CreateAttachmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create attachment
	attachment := models.Attachment{
		FilePath: input.FilePath,
		FileType: input.FileType,
	}
	db.DB.Create(&attachment)
	c.JSON(http.StatusCreated, gin.H{"data": attachment})
}
