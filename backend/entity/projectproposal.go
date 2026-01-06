package entity

import (
	"gorm.io/gorm"
	"time"
)

type ProjectProposal struct {
	gorm.Model
	Projectname 	string `validate:"required~Project name is required, min=5,max=100~Project name must be between 5 to 100 characters"`
	Description 	string `validate:"max=500~Description must be at least 500 characters"`
	StartDate   	time.Time `validate:"required~Start date is required"`
	DueDate     	time.Time `validate:"required~Due date is required,gtfield=StartDate~Due date must be after start date"`
	Status      	string `validate:"required~Status is required,oneof=pending approved rejected~Status must be either pending, approved, or rejected"`
	RequstedAmount 	float64 `validate:"required~Requested amount is required,range(1|999999)~Requested amount must be between 1 and 999,999"`
	CoverPath   	string `validate:"required~Cover path is required"` 
}