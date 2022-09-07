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

// get exercise by id
type ExerciseGetByIdRequest struct {
	ExerciseId int
}

type ExerciseGetByIdResponse struct {
	ExerciseModel
	Questions []*QuestionModel `json:"questions"`
}

type ExerciseGetScoreRequest struct {
	ExerciseId int
}

type ExerciseGetScoreResponse struct {
	Score int `json:"score"`
}
