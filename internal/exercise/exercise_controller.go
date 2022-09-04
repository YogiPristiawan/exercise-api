package exercise

import (
	"exercise-api/internal/exercise/entities"
	"exercise-api/internal/presentation"

	"github.com/gin-gonic/gin"
)

type exerciseController struct {
	exerciseService ExerciseService
}

func NewExerciseController(
	exerciseService ExerciseService,
) *exerciseController {
	return &exerciseController{
		exerciseService: exerciseService,
	}
}

func (e *exerciseController) Create(c *gin.Context) {
	in := entities.ExerciseCreateRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := e.exerciseService.Create(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)

}
