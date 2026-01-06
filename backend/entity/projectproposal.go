package entity

import (
	"gorm.io/gorm"
	"time"
)

type ProjectProposal struct {
	gorm.Model
	Projectname 	string
	Description 	string
	StartDate   	time.Time
	DueDate     	time.Time
	Status      	string
	RequstedAmount 	float64
}