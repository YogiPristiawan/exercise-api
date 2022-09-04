package entities

import "exercise-api/internal/shared/entities"

type ExerciseModel struct {
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	AuthorId    int
}

func (e *ExerciseModel) TableName() string {
	return "exercises"
}

// create exercise
type ExerciseCreateRequest struct {
	entities.RequestMetaData
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type ExerciseCreateResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
