package repositories

import (
	"exercise-api/internal/exercise/entities"
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/exercise/model"

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

func (q *questionRepository) GetById(questionId int) (question *model.GetQuestionById, err error) {
	err = q.db.Table("questions").
		Select(
			"id",
			"exercise_id",
			"body",
		).
		Where("id = ?", questionId).First(&question).Error
	return
}

func (q *questionRepository) FindByExerciseId(exerciseId int) (questions []*entities.QuestionModel, err error) {
	err = q.db.Where("exercise_id = ?", exerciseId).Find(&questions).Error
	return
}
