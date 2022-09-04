package services

import "exercise-api/internal/exercise/entities"

type ExerciseValidator interface {
	ValidateCreateExercisePayload(*entities.ExerciseCreateRequest) error
}

type ExerciseRepository interface {
	Create(*entities.ExerciseModel) error
}
