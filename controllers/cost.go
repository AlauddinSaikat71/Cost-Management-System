package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"log"
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
		log.Fatal(err)
	}
	toDate, err := time.Parse("2006-01-02", toDateString)
	if err != nil {
		log.Fatal(err)
	}

	// query in costs table according to date
	var costs []models.Cost
	if err := db.CostDbContext.Where("created_at>=? AND created_at<=?", fromDate, toDate).Find(&costs).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": costs})
}

//PATCH /costs/:id
//Update  a cost
func UpdateCost(c *gin.Context) {
	// Validate input
	var input dtos.UpdateCostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var cost models.Cost
	if err := db.CostDbContext.Where("id = ?", c.Param("id")).First(&cost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Model(&cost).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": cost})
}

//DELETE costs/:id
// delete cost by id
func DeleteCost(c *gin.Context) {
	//get model if exist
	var cost models.Cost
	if err := db.CostDbContext.Where("id = ?", c.Param("id")).First(&cost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&cost)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
