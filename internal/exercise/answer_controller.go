package exercise

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/presentation"
	"exercise-api/internal/shared/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type answerController struct {
	answerService AnswerService
}

func NewAnswerController(
	answerService AnswerService,
) *answerController {
	return &answerController{
		answerService: answerService,
	}
}

func (a *answerController) Create(c *gin.Context) {
	// collect params
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	if err != nil {
		out := entities.CommonResult{
			ResCode:    400,
			ResMessage: "parameter exercise id harus berupa angka",
		}
		presentation.WriteRestOut(c, out, &out)
		return
	}

	questionId, err := strconv.Atoi(c.Param("questionId"))
	if err != nil {
		out := entities.CommonResult{
			ResCode:    400,
			ResMessage: "parameter question id harus berupa angka",
		}
		presentation.WriteRestOut(c, out, &out)
		return
	}

	in := exerciseEntities.AnswerCreateRequest{
		ExerciseId: exerciseId,
		QuestionId: questionId,
	}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := a.answerService.Create(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}
