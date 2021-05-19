package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"fmt"
	"net/http"
	"time"

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

//GET /costs
// find costs between from-date and to-date
func FindCosts(c *gin.Context) {
	fromDateString := c.Param("from")
	//fmt.Printf(fromDateString)
	toDateString := c.Param("to")
	//fmt.Printf(toDateString)

	// parsing date object from string
	fromDate, err := time.Parse("2006-01-02", fromDateString)
	if err != nil {
		panic(err)
	}
	fmt.Println(fromDate)
	toDate, err := time.Parse("2006-01-02", toDateString)
	if err != nil {
		panic(err)
	}

	var costs []models.Cost
	if err := db.CostDbContext.Where("created_at BETWEEN ? AND ?", fromDate, toDate).Find(&costs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": costs})
}
