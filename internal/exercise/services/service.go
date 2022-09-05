package services

import (
	"exercise-api/internal/exercise/entities"
	"exercise-api/internal/exercise/shared"
)

// exercise
type ExerciseValidator interface {
	ValidateCreateExercisePayload(*entities.ExerciseCreateRequest) error
}

type ExerciseRepository interface {
	Create(*entities.ExerciseModel) error
	GetById(int) (*shared.GetExerciseByIdDTO, error)
}

// question
type QuestionValidator interface {
	ValidateCreateQuestionPayload(*entities.QuestionCreateRequest) error
}

type QuestionRepository interface {
	Create(*entities.QuestionModel) error
	GetById(int) (*shared.GetQuestionByIdDTO, error)
}

// answer
type AnswerValidator interface {
	ValidateCreateAnswerPayload(*entities.AnswerCreateRequest) error
}

type AnswerRepository interface {
	Create(*entities.AnswerModel) error
}
