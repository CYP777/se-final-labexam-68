package service

import (
	"errors" // เพิ่ม import errors
	"se-lab-exam-final/entity"
	"github.com/asaskevich/govalidator"
)

func ValidateProjectProposal(proposal *entity.ProjectProposal) (bool, error) {
	result, err := govalidator.ValidateStruct(proposal)
	if !result {
		return result, err
	}
	if !proposal.DueDate.After(proposal.StartDate) {
		return false, errors.New("Due date must be after start date")
	}
	return true, nil
}