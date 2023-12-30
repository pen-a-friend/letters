package database

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"-"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Users struct{
	BaseModel
	Email string 
	FirstName string
	LastName string
	ClerkId string
}

type Mail struct{
	BaseModel
	To string
	From string
	Subject string
	Body string
	Status string
	Message string
}