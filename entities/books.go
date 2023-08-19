package entities

import (
	"time"

	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type Book struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string    `validate:"required" json:"title" gorm:"type:string;size:256;not null"`
	Author      string    `validate:"required" json:"author" gorm:"type:string;size:256;not null"`
	Description string    `validate:"required" json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
