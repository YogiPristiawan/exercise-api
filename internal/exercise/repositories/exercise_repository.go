package repositories

import (
	"exercise-api/internal/exercise/entities"
	"log"

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
	res := e.db.Create(exercise)
	log.Print("hasil dari exercise", exercise)
	return res.Error
}
