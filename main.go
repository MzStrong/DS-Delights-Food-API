package main

import (
	"github.com/MzStrong/DS-Delights-Food-API/database"
	"github.com/MzStrong/DS-Delights-Food-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	database.ConnectDB()

	// Seed ข้อมูลครั้งแรก
	database.SeedDB()

	// ตั้งค่า Routes
	r := gin.Default()
	routes.SetupRoutes(r)

	// รันเซิร์ฟเวอร์
	r.Run(":8080")
}