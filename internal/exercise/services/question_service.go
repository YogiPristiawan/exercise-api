package services

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/shared/entities"
	"fmt"
)

type questionService struct {
	validator          QuestionValidator
	questionRepository QuestionRepository
	exerciseRepository ExerciseRepository
}

func NewQuestionService(
	validator QuestionValidator,
	questionRepository QuestionRepository,
	exerciseRepository ExerciseRepository,
) *questionService {
	return &questionService{
		validator:          validator,
		questionRepository: questionRepository,
		exerciseRepository: exerciseRepository,
	}
}

func (q *questionService) Create(in *exerciseEntities.QuestionCreateRequest) (out entities.BaseResponse[exerciseEntities.QuestionCreateResponse]) {
	// validate payload
	if err := q.validator.ValidateCreateQuestionPayload(in); err != nil {
		out.SetCode(400, err)
		return
	}

	// validate if exerciseID exists
	_, err := q.exerciseRepository.GetById(in.ExerciseId)
	switch castDatabaseError(err) {
	case 404:
		out.SetCode(404, fmt.Errorf("invalid exercise id"))
		return
	case 500:
		out.SetCode(500, err)
		return
	}

	// create question
	question := &exerciseEntities.QuestionModel{
		ExerciseId:    in.ExerciseId,
		Body:          in.Body,
		OptionA:       in.OptionA,
		OptionB:       in.OptionB,
		OptionC:       in.OptionC,
		OptionD:       in.OptionD,
		CorrectAnswer: in.CorrectAnswer,
		Score:         in.Score,
	}
	err = q.questionRepository.Create(question)
	switch castDatabaseError(err) {
	case 500:
		out.SetCode(500, err)
		return
	}

	out.Message = "question created"
	out.Data.Id = question.ID
	out.Data.ExerciseId = question.ExerciseId
	out.SetCode(201, nil)
	return
}
