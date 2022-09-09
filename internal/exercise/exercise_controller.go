package exercise

import (
	exerciseEntities "exercise-api/internal/exercise/entities"
	"exercise-api/internal/presentation"
	"exercise-api/internal/shared/entities"
	"strconv"

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
	authUserId, _ := c.Get("user_id")
	in := exerciseEntities.ExerciseCreateRequest{
		RequestMetaData: entities.RequestMetaData{
			AuthUserId: int(authUserId.(float64)),
		},
	}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := e.exerciseService.Create(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)

}

func (e *exerciseController) GetScore(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	authUserId, _ := c.Get("user_id")
	if err != nil {
		out := entities.CommonResult{
			ResCode:    400,
			ResMessage: "parameter harus berupa angka",
		}
		presentation.WriteRestOut(c, out, &out)
		return
	}

	in := exerciseEntities.ExerciseGetScoreRequest{
		ExerciseId: exerciseId,
		RequestMetaData: entities.RequestMetaData{
			AuthUserId: int(authUserId.(float64)),
		},
	}

	out := e.exerciseService.GetScore(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (e *exerciseController) GetById(c *gin.Context) {
	exerciseId, err := strconv.Atoi(c.Param("exerciseId"))
	if err != nil {
		out := entities.CommonResult{
			ResCode:    400,
			ResMessage: "parameter harus berupa angka",
		}
		presentation.WriteRestOut(c, out, &out)
		return
	}

	in := exerciseEntities.ExerciseGetByIdRequest{
		ExerciseId: exerciseId,
	}

	out := e.exerciseService.GetById(&in)
	presentation.WriteRestOut(c, out, &out.CommonResult)
}
