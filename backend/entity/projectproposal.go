package entity

import (
	"gorm.io/gorm"
	"time"
)

type ProjectProposal struct {
	gorm.Model
	Projectname 	string `valid:"required~Project name is required, stringlength(5|100)~Project name must be between 5 to 100 characters"`
	Description 	string `valid:"stringlength(0|500)~Description must be less than 500 characters"`
	StartDate   	time.Time `valid:"required~Start date is required"`
	DueDate     	time.Time `valid:"required~Due date is required"`
	Status      	string `valid:"required~Status is required,in(pending|approved|rejected)~Status must be either pending or approved or rejected"`
	RequstedAmount 	float64 `valid:"required~Requested amount is required,range(1|999999)~Requested amount must be between 1 and 999999"`
	CoverPath   	string `valid:"required~Cover path is required"` 
}