package validator

import (
	"exercise-api/internal/account/entities"
	sharedValidator "exercise-api/internal/shared/validator"

	"github.com/go-playground/validator/v10"
)

var (
	customErrorMessage = sharedValidator.CustomErrorMessage
)

type accountValidator struct {
	validator *validator.Validate
}

func NewAccountValidator(
	validator *validator.Validate,
) *accountValidator {
	return &accountValidator{
		validator: validator,
	}
}

func (a *accountValidator) ValidateRegisterPayload(in *entities.RegisterRequest) (err error) {
	err = a.validator.Struct(in)
	err = customErrorMessage(err)
	return
}

func (a *accountValidator) ValidateLoginPayload(in *entities.LoginRequest) (err error) {
	err = a.validator.Struct(in)
	err = customErrorMessage(err)
	return
}
