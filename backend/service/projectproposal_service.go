package service

import (
	"se-lab-exam-final/entity"
	"github.com/asaskevich/govalidator"
)

func ValidateProjectProposal(proposal *entity.ProjectProposal) (bool, error) {
	return govalidator.ValidateStruct(proposal)
}