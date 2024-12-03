package models

type Menu struct {
	ID       uint      `gorm:"primaryKey"`
	Name     string    `json:"name" binding:"required"`
	Diseases []Disease `gorm:"many2many:menu_diseases;"` // ความสัมพันธ์ Many-to-Many
}

type Disease struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Menus []Menu `gorm:"many2many:menu_diseases;"` // ความสัมพันธ์ Many-to-Many
}

type MenuDisease struct {
	MenuID    uint `gorm:"primaryKey"`
	DiseaseID uint `gorm:"primaryKey"`
}
