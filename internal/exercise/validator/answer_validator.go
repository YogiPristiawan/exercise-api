package validator

import (
	"exercise-api/internal/exercise/entities"

	"github.com/go-playground/validator/v10"
)

type answerValidator struct {
	validator *validator.Validate
}

func NewAnswerValidator(
	validator *validator.Validate,
) *answerValidator {
	return &answerValidator{
		validator: validator,
	}
}

func (a *answerValidator) ValidateCreateAnswerPayload(in *entities.AnswerCreateRequest) error {
	err := a.validator.Struct(in)
	return customErroMessage(err)
}
