package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name  string `json:"name" binding:"required" gorm:"not null"`
	Age   string `json:"age" binding:"required" gorm:"not null"`
	Books []Book `gorm:"foreignKey:AuthorRefer"`
}
