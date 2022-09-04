package validator

import (
	"exercise-api/internal/exercise/entities"

	"github.com/go-playground/validator/v10"
)

type questionValidator struct {
	validator *validator.Validate
}

func NewQuestionValidator(
	validator *validator.Validate,
) *questionValidator {
	return &questionValidator{
		validator: validator,
	}
}

func (q *questionValidator) ValidateCreateQuestionPayload(in *entities.QuestionCreateRequest) (err error) {
	err = q.validator.Struct(in)
	err = customErroMessage(err)
	return
}
