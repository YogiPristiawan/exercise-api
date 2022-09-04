package repositories

import (
	exerciseEntities "exercise-api/internal/exercise/entities"

	"gorm.io/gorm"
)

type questionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(
	db *gorm.DB,
) *questionRepository {
	return &questionRepository{
		db: db,
	}
}

func (q *questionRepository) Create(question *exerciseEntities.QuestionModel) (err error) {
	err = q.db.Create(&question).Error
	return
}
