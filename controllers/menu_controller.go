package controllers

import (
	"github.com/MzStrong/DS-Delights-Food-API/database"
	"github.com/MzStrong/DS-Delights-Food-API/models"
	"github.com/gin-gonic/gin"
)

func CreateMenu(c *gin.Context) {
	var request struct {
		Name      string  `json:"name" binding:"required"`
		DiseaseIDs []uint `json:"disease_ids"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var diseases []models.Disease
	if len(request.DiseaseIDs) > 0 {
		database.DB.Where("id IN ?", request.DiseaseIDs).Find(&diseases)
	}

	menu := models.Menu{Name: request.Name, Diseases: diseases}
	if err := database.DB.Create(&menu).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Menu created successfully", "menu": menu})
}

func GetMenus(c *gin.Context) {
	var menus []models.Menu
	database.DB.Preload("Diseases").Find(&menus)
	c.JSON(200, menus)
}
