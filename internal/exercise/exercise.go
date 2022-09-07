package exercise

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/shared/entities"
)

type ExerciseService interface {
	Create(*exerciseEntities.ExerciseCreateRequest) entities.BaseResponse[exerciseEntities.ExerciseCreateResponse]
	GetScore(*exerciseEntities.ExerciseGetScoreRequest) entities.BaseResponse[exerciseEntities.ExerciseGetScoreResponse]
}

type QuestionService interface {
	Create(*exerciseEntities.QuestionCreateRequest) entities.BaseResponse[exerciseEntities.QuestionCreateResponse]
}

type AnswerService interface {
	Create(*exerciseEntities.AnswerCreateRequest) entities.BaseResponse[exerciseEntities.AnswerCreateResponse]
}
