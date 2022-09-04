package entities

import "exercise-api/internal/shared/entities"

type QuestionModel struct {
	ID            int `gorm:"primaryKey"`
	ExerciseId    int
	Body          string
	OptionA       string
	OptionB       string
	OptionC       string
	OptionD       string
	CorrectAnswer string
	Score         int
	AuthorId      int
	CreatedAt     int `gorm:"autoCreateTime"`
	UpdatedAt     int `gorm:"autoUpdateTime"`
}

func (q *QuestionModel) TableName() string {
	return "questions"
}

type QuestionCreateRequest struct {
	entities.RequestMetaData
	ExerciseId    int
	Body          string `json:"body" validate:"required"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	Score         int    `json:"score" validate:"required"`
	CorrectAnswer string `json:"correct_answer" validate:"required,max=1,oneof=a b c d"`
}

type QuestionCreateResponse struct {
	Id         int `json:"id"`
	ExerciseId int `json:"exercise_id"`
}
