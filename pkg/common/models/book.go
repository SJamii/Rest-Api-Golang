package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	AuthorRefer uuid.UUID
}
