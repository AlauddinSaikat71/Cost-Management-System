package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /costattachment
// create new costattachment
func CreateCostAttachment(c *gin.Context) {
	//validate input
	var input dtos.CreateCostAttachmentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create costattachment
	costattachment := models.CostAttachment{
		CostId: input.CostId,
		AttachmentId: input.AttachmentId,
		CreatedBy: input.CreatedBy,	
	}
	db.DB.Create(&costattachment)
	c.JSON(http.StatusCreated, gin.H{"data": costattachment})
}