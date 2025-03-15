package model

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name string    `json:"category_name"`
}

type ProductCategory struct {
	CategoryId   uuid.UUID `json:"product_category_id" gorm:"type:uuid;primary_key"`
	CategoryName string    `json:"product_category_name" gorm:"type:varchar(255);not null"`
	Status       string    `json:"status" gorm:"type:varchar(25);not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"type:timestamp;not null"`
}
