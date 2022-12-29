package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	ID    uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name  string    `json:"name" binding:"required" gorm:"not null"`
	Age   string    `json:"age" binding:"required" gorm:"not null"`
	Books []Book    `gorm:"foreignKey:AuthorRefer"`
}
