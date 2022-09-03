package validator

import (
	"exercise-api/internal/account/entities"
	sharedValidator "exercise-api/internal/shared/validator"

	"github.com/go-playground/validator/v10"
)

var (
	customErrorMessage = sharedValidator.CustomErrorMessage
)

type studentValidator struct {
	validator *validator.Validate
}

func NewAccountValidator(
	validator *validator.Validate,
) *studentValidator {
	return &studentValidator{
		validator: validator,
	}
}

func (s *studentValidator) ValidateRegisterPayload(in *entities.RegisterRequest) (err error) {
	err = s.validator.Struct(in)
	err = customErrorMessage(err)
	return
}
