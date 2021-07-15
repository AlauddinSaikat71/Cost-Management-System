package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /payments
// create new payment
func CreatePayment(c *gin.Context) {
	//validate input
	var input dtos.CreatePaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//create payment
	payment := models.Payment{
		Method:    input.Method,
		Amount:    input.Amount,
		CreatedBy: input.CreatedBy,
		PaidBy:    input.PaidBy,
		Meta:      input.Meta,
	}
	db.DB.Create(&payment)
	c.JSON(http.StatusCreated, gin.H{"data": payment})
}

//PATCH /payments/:id
//Update  a payment
func UpdatePayment(c *gin.Context) {
	// Validate input
	var input dtos.UpdatePaymentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var payment models.Payment
	if err := db.PaymentDbContext.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Model(&payment).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": payment})
}

//GET /payments/:id
//get payment by id
func FindPayment(c *gin.Context) {
	//get model if exist
	var payment models.Payment
	if err := db.PaymentDbContext.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": payment})
}

//GET /payments
//GET all payments
func FindPayments(c *gin.Context) {
	var payments []models.Payment
	db.DB.Find(&payments)

	c.JSON(http.StatusOK, gin.H{"data": payments})
}

//DELETE payments/:id
// delete payment by id
func DeletePayment(c *gin.Context) {
	//get model if exist
	var payment models.Payment
	if err := db.PaymentDbContext.Where("id = ?", c.Param("id")).First(&payment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&payment)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
