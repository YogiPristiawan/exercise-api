package routes

import "github.com/gin-gonic/gin"

type ExerciseController interface {
	Create(c *gin.Context)
}

func NewExerciseRoutes(r *gin.Engine, exerciseController ExerciseController) {
	g := r.Group("exercises")
	{
		g.POST("", exerciseController.Create)
	}
}
