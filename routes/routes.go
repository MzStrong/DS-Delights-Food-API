package routes

import (
	"github.com/MzStrong/DS-Delights-Food-API/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/menus", controllers.CreateMenu)
	r.GET("/menus", controllers.GetMenus)

	r.POST("/diseases", controllers.CreateDisease)
	r.GET("/diseases", controllers.GetDiseases)
}
