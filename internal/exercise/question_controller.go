package exercise

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/presentation"
	"exercise-api/internal/shared/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type questionController struct {
	questionService QuestionService
}

func NewQuestionController(
	questionService QuestionService,
) *questionController {
	return &questionController{
		questionService: questionService,
	}
}

func (q *questionController) Create(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	if err != nil {
		out := entities.CommonResult{
			ResCode:    400,
			ResMessage: "parameter harus berupa angka",
		}
		presentation.WriteRestOut(c, out, &out)
		return
	}

	in := exerciseEntities.QuestionCreateRequest{
		ExerciseId: exerciseId,
	}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := q.questionService.Create(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}
