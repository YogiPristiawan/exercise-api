package validator

import (
	"exercise-api/internal/exercise/entities"
	sharedValidator "exercise-api/internal/shared/validator"

	"github.com/go-playground/validator/v10"
)

var (
	customErroMessage = sharedValidator.CustomErrorMessage
)

type exerciseValidator struct {
	validator *validator.Validate
}

func NewExerciseValidator(
	validator *validator.Validate,
) *exerciseValidator {
	return &exerciseValidator{
		validator: validator,
	}
}

func (e *exerciseValidator) ValidateCreateExercisePayload(in *entities.ExerciseCreateRequest) (err error) {
	err = e.validator.Struct(in)
	err = customErroMessage(err)
	return
}
