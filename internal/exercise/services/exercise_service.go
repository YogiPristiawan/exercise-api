package services

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/shared/databases"
	"exercise-api/internal/shared/entities"
	"fmt"
	"strings"
	"sync"
)

var (
	castDatabaseError = databases.CastDatabaseError
)

type exerciseService struct {
	validator          ExerciseValidator
	exerciseRepository ExerciseRepository
	questionRepository QuestionRepository
}

func NewExerciseService(
	validator ExerciseValidator,
	exerciseRepository ExerciseRepository,
	questionRepository QuestionRepository,
) *exerciseService {
	return &exerciseService{
		validator:          validator,
		exerciseRepository: exerciseRepository,
		questionRepository: questionRepository,
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

func (e *exerciseService) GetScore(in *exerciseEntities.ExerciseGetScoreRequest) (out entities.BaseResponse[exerciseEntities.ExerciseGetScoreResponse]) {
	// verify if exercise exists
	_, err := e.exerciseRepository.GetById(in.ExerciseId)
	switch castDatabaseError(err) {
	case 404:
		out.SetCode(404, fmt.Errorf("exercise tidak ditemukan"))
		return
	case 500:
		out.SetCode(500, err)
		return
	}

	exercises, err := e.exerciseRepository.FindUserQuestionAnswer(in.ExerciseId, in.AuthUserId)
	switch castDatabaseError(err) {
	case 404:
		out.SetCode(404, fmt.Errorf("exercise tidak ditemukan"))
		return
	case 500:
		out.SetCode(500, err)
		return
	}

	// calculate score
	var wg sync.WaitGroup
	var m sync.Mutex
	var score int

	for _, val := range exercises {
		wg.Add(1)
		go func(val map[string]interface{}) {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()

			if strings.EqualFold(val["correct_answer"].(string), val["user_answer"].(string)) {
				score += int(val["score"].(int32))
			}
		}(val)
	}
	wg.Wait()

	out.Message = "score"
	out.Data.Score = score
	return
}

func (e *exerciseService) GetById(in *exerciseEntities.ExerciseGetByIdRequest) (out entities.BaseResponse[exerciseEntities.ExerciseGetByIdResponse]) {
	// get exercise
	exercise, err := e.exerciseRepository.GetById(in.ExerciseId)
	switch castDatabaseError(err) {
	case 404:
		out.SetCode(404, fmt.Errorf("exercise tidak ditemukan"))
		return
	case 500:
		out.SetCode(500, err)
		return
	}

	// find questions
	questions, err := e.questionRepository.FindByExerciseId(exercise.Id)
	switch castDatabaseError(err) {
	case 500:
		out.SetCode(500, err)
		return
	}

	out.Message = "exercises and questions"
	out.Data.Id = exercise.Id
	out.Data.Title = exercise.Title
	out.Data.Description = exercise.Description
	out.Data.AuthorId = exercise.AuthorId
	out.Data.Questions = []map[string]interface{}{}

	if len(questions) > 0 {
		for _, val := range questions {
			question := make(map[string]interface{})
			question["id"] = val.ID
			question["body"] = val.Body
			question["option_a"] = val.OptionA
			question["option_b"] = val.OptionB
			question["option_c"] = val.OptionC
			question["option_d"] = val.OptionD
			question["score"] = val.Score
			question["created_at"] = val.CreatedAt
			question["updated_at"] = val.UpdatedAt

			out.Data.Questions = append(out.Data.Questions, question)
		}
	}

	return
}
