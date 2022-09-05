package entities

import "exercise-api/internal/shared/entities"

type AnswerModel struct {
	ID         int `gorm:"primaryKey"`
	ExerciseId int
	QuestionId int
	UserId     int
	Answer     string
	CreatedAt  int `gorm:"autoCreateTime"`
	UpdatedAt  int `gorm:"autoUpdateTime"`
}

func (a *AnswerModel) TableName() string {
	return "answers"
}

type AnswerCreateRequest struct {
	entities.RequestMetaData
	ExerciseId int
	QuestionId int
	Answer     string `json:"answer" validate:"required,max=1,oneof=a b c d"`
}

type AnswerCreateResponse struct {
	Id int `json:"id"`
}
