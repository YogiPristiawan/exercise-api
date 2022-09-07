package services

import (
	answerEntities "exercise-api/internal/exercise/entities"
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/shared/entities"
	"fmt"
	"log"
)

type answerService struct {
	validator          AnswerValidator
	questionRepository QuestionRepository
	answerRepository   AnswerRepository
}

func NewAnswerService(
	answerValidator AnswerValidator,
	questionRepository QuestionRepository,
	answerRepository AnswerRepository,
) *answerService {
	return &answerService{
		validator:          answerValidator,
		questionRepository: questionRepository,
		answerRepository:   answerRepository,
	}
}

func (a *answerService) Create(in *answerEntities.AnswerCreateRequest) (out entities.BaseResponse[answerEntities.AnswerCreateResponse]) {
	// validate payload
	if err := a.validator.ValidateCreateAnswerPayload(in); err != nil {
		out.SetCode(400, err)
		return
	}

	// check if exercise and questions exists
	question, err := a.questionRepository.GetById(in.QuestionId)
	log.Print(question)
	switch castDatabaseError(err) {
	case 404:
		out.SetCode(404, fmt.Errorf("data question tidak ditemukan"))
		return
	case 500:
		out.SetCode(400, err)
		return
	}

	if question.ExerciseId != in.ExerciseId {
		out.SetCode(400, fmt.Errorf("data exercise tidak ditemukan"))
		return
	}

	// create answer
	answer := exerciseEntities.AnswerModel{
		ExerciseId: in.ExerciseId,
		QuestionId: in.QuestionId,
		UserId:     in.RequestMetaData.AuthUserId,
		Answer:     in.Answer,
	}
	err = a.answerRepository.Create(&answer)
	switch castDatabaseError(err) {
	case 500:
		out.SetCode(500, err)
		return
	}

	out.Message = "answer created"
	out.Data.Id = answer.ID
	out.SetCode(201, nil)
	return
}
