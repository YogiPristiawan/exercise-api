package repositories

import (
	"exercise-api/internal/exercise/entities"
	"exercise-api/internal/exercise/model"

	"gorm.io/gorm"
)

type exerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(
	db *gorm.DB,
) *exerciseRepository {
	return &exerciseRepository{
		db: db,
	}
}

func (e *exerciseRepository) Create(exercise *entities.ExerciseModel) (err error) {
	err = e.db.Create(exercise).Error
	return
}

func (e *exerciseRepository) GetById(id int) (exercise *model.GetExerciseById, err error) {
	err = e.db.
		Table("exercises").
		Select(
			"id",
			"title",
			"description",
			"author_id",
		).
		Where("id = ?", id).First(&exercise).Error
	return
}

func (e *exerciseRepository) FindUserQuestionAnswer(exerciseId int, userId int) (results []map[string]interface{}, err error) {
	err = e.db.Raw(`
		SELECT
			questions.correct_answer,
			questions.score,
			answers.answer AS user_answer
		FROM
			answers
			INNER JOIN questions ON questions.id = answers.question_id AND questions.exercise_id = ?
		WHERE
			answers.user_id = ?
	`, exerciseId, userId).Find(&results).Error
	return
}
