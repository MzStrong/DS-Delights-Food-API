package database

import (
	"github.com/MzStrong/DS-Delights-Food-API/models"
	"log"
)

func SeedDB() {
	// ตรวจสอบว่ามีโรคในฐานข้อมูลหรือยัง
	var count int64
	DB.Model(&models.Disease{}).Count(&count)
	if count > 0 {
		log.Println("Database already seeded")
		return
	}

	// เพิ่มโรคตัวอย่าง (ภาษาไทย)
	diseases := []models.Disease{
		{Name: "เบาหวาน"},
		{Name: "ความดันโลหิตสูง"},
		{Name: "โรคหัวใจ"},
		{Name: "ไตวาย"},
		{Name: "ไขมันในเลือดสูง"},
	}
	if err := DB.Create(&diseases).Error; err != nil {
		log.Fatalf("Failed to seed diseases: %v", err)
	}

	// เพิ่มเมนูตัวอย่าง (ภาษาไทย)
	menus := []models.Menu{
		{Name: "ข้าวมันไก่"},         // มีโรค
		{Name: "ยำผักรวม"},         // มีโรค
		{Name: "แกงเขียวหวานไก่"},  // มีโรค
		{Name: "ปลานึ่งมะนาว"},      // มีโรค
		{Name: "ลาบหมู"},           // มีโรค
		{Name: "ไข่พะโล้"},         // มีโรค
		{Name: "ส้มตำไทย"},         // มีโรค
		{Name: "ต้มยำกุ้ง"},         // มีโรค
		{Name: "ข้าวผัดกุ้ง"},       // ไม่มีโรค
		{Name: "แกงส้มปลาช่อน"},     // ไม่มีโรค
	}
	if err := DB.Create(&menus).Error; err != nil {
		log.Fatalf("Failed to seed menus: %v", err)
	}

	// เพิ่มความสัมพันธ์ระหว่างเมนูกับโรค
	menuDiseases := []models.MenuDisease{
		{MenuID: 1, DiseaseID: 1}, // ข้าวมันไก่ ไม่เหมาะสำหรับ เบาหวาน
		{MenuID: 1, DiseaseID: 5}, // ข้าวมันไก่ ไม่เหมาะสำหรับ ไขมันในเลือดสูง
		{MenuID: 2, DiseaseID: 2}, // ยำผักรวม ไม่เหมาะสำหรับ ความดันโลหิตสูง
		{MenuID: 3, DiseaseID: 1}, // แกงเขียวหวานไก่ ไม่เหมาะสำหรับ เบาหวาน
		{MenuID: 3, DiseaseID: 4}, // แกงเขียวหวานไก่ ไม่เหมาะสำหรับ ไตวาย
		{MenuID: 4, DiseaseID: 3}, // ปลานึ่งมะนาว ไม่เหมาะสำหรับ โรคหัวใจ
		{MenuID: 5, DiseaseID: 2}, // ลาบหมู ไม่เหมาะสำหรับ ความดันโลหิตสูง
		{MenuID: 6, DiseaseID: 4}, // ไข่พะโล้ ไม่เหมาะสำหรับ ไตวาย
		{MenuID: 6, DiseaseID: 5}, // ไข่พะโล้ ไม่เหมาะสำหรับ ไขมันในเลือดสูง
		{MenuID: 7, DiseaseID: 1}, // ส้มตำไทย ไม่เหมาะสำหรับ เบาหวาน
		{MenuID: 8, DiseaseID: 3}, // ต้มยำกุ้ง ไม่เหมาะสำหรับ โรคหัวใจ
	}
	if err := DB.Create(&menuDiseases).Error; err != nil {
		log.Fatalf("Failed to seed DS-Delights-Food-API relations: %v", err)
	}

	log.Println("Database seeded successfully")
}
