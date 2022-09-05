package repositories

import (
	"exercise-api/internal/exercise/entities"

	"gorm.io/gorm"
)

type answerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *answerRepository {
	return &answerRepository{
		db: db,
	}
}

func (a *answerRepository) Create(answer *entities.AnswerModel) error {
	return a.db.Create(&answer).Error
}
