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
		).
		Where("id = ?", id).First(&exercise).Error
	return
}
