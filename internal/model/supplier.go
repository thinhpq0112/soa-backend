package model

import "github.com/google/uuid"

type Supplier struct {
	Id   uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name string    `json:"name" gorm:"type:varchar(255);not null;unique"`
}
