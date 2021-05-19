package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /costs
// create new cost
func CreateCost(c *gin.Context) {
	//validate input
	var input dtos.CreateCostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create cost
	cost := models.Cost{
		Title:       input.Title,
		Description: input.Description,
		Amount:      input.Amount,
		Payment_Id:  input.Payment_Id,
	}
	db.DB.Create(&cost)
	c.JSON(http.StatusCreated, gin.H{"data": cost})
}
