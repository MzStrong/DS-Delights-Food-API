package controllers

import (
	"github.com/MzStrong/DS-Delights-Food-API/database"
	"github.com/MzStrong/DS-Delights-Food-API/models"
	"github.com/gin-gonic/gin"
)

func CreateDisease(c *gin.Context) {
	var disease models.Disease
	if err := c.ShouldBindJSON(&disease); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&disease).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Disease created successfully", "disease": disease})
}

func GetDiseases(c *gin.Context) {
	var diseases []models.Disease
	database.DB.Preload("Menus").Find(&diseases)
	c.JSON(200, diseases)
}
