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

func (q *answerRepository) VerifyExistsAnswer(userId, exerciseId, questionId int) (int64, error) {
	var count int64
	err := q.db.
		Model(&entities.AnswerModel{}).
		Where("user_id = ?", userId).
		Where("exercise_id = ?", exerciseId).
		Where("question_id = ?", questionId).Count(&count).Error
	return count, err
}
