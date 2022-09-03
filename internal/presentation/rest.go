package presentation

import (
	"exercise-api/internal/shared/entities"

	"github.com/gin-gonic/gin"
)

func ReadRestIn[T interface{}](c *gin.Context, in *T) bool {
	if err := c.ShouldBindJSON(in); err != nil {
		c.AbortWithStatusJSON(400, struct {
			Message string `json:"message"`
		}{
			Message: "invalid request body",
		})
		return false
	}
	return true
}

func WriteRestOut[T interface{}](c *gin.Context, out T, cr *entities.CommonResult) {
	if cr.ResCode == 0 {
		c.JSON(200, out)
		return
	}

	if cr.ResCode < 400 {
		c.JSON(cr.ResCode, out)
		return
	}

	if cr.ResCode >= 400 {
		c.AbortWithStatusJSON(cr.ResCode, struct {
			Message string `json:"message"`
		}{
			Message: cr.ResMessage,
		})
		return
	}
}
