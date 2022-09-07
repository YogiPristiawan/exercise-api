package services

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/shared/databases"
	"exercise-api/internal/shared/entities"
)

var (
	castDatabaseError = databases.CastDatabaseError
)

type exerciseService struct {
	validator          ExerciseValidator
	exerciseRepository ExerciseRepository
}

func NewExerciseService(
	validator ExerciseValidator,
	exerciseRepository ExerciseRepository,
) *exerciseService {
	return &exerciseService{
		validator:          validator,
		exerciseRepository: exerciseRepository,
	}
}

func (e *exerciseService) Create(in *exerciseEntities.ExerciseCreateRequest) (out entities.BaseResponse[exerciseEntities.ExerciseCreateResponse]) {
	// valdiate paylaod
	if err := e.validator.ValidateCreateExercisePayload(in); err != nil {
		out.SetCode(400, err)
		return
	}

	// create exericise
	exercise := &exerciseEntities.ExerciseModel{
		Title:       in.Title,
		Description: in.Description,
		AuthorId:    in.RequestMetaData.AuthUserId,
	}
	if err := e.exerciseRepository.Create(exercise); err != nil {
		switch castDatabaseError(err) {
		case 500:
			out.SetCode(500, err)
			return
		}
	}

	out.Message = "exercise created"
	out.Data.Id = exercise.Id
	out.Data.Title = exercise.Title
	out.SetCode(201, nil)
	return
}

func (e *exerciseService) GetScore(in *exerciseEntities.ExerciseGetByIdRequest) (out entities.BaseResponse[exerciseEntities.ExerciseGetByIdResponse]) {
	// get exercise
	exercise := 
}
