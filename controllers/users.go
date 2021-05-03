package controllers

import (
	"costmanagement/db"
	"costmanagement/models"
	"costmanagement/models/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

//POST /users
// create new user
func CreateUser(c *gin.Context) {
	//validate input
	var input dtos.CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//create user
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	}
	db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

//PATCH /users/:id
//Update  a user
func UpdateUser(c *gin.Context) {
	// Validate input
	var input dtos.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var user models.User
	if err := db.UserDbContext.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

//GET /users/:id
//get user by id
func FindUser(c *gin.Context) {
	//get model if exist
	var user models.User
	if err := db.UserDbContext.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

//GET /users
//GET all users
func FindUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

//DELETE users/:id
// delete userby id
func DeleteUser(c *gin.Context) {
	//get model if exist
	var user models.User
	if err := db.UserDbContext.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	db.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
