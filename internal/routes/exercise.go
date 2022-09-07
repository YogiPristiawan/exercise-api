package routes

import (
	"exercise-api/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

type ExerciseController interface {
	Create(c *gin.Context)
	GetScore(c *gin.Context)
}

type QuestionController interface {
	Create(c *gin.Context)
}

type AnswerController interface {
	Create(c *gin.Context)
}

func NewExerciseRoutes(
	r *gin.Engine, exerciseController ExerciseController,
	questionController QuestionController,
	answerController AnswerController,
	jwtMiddleware gin.HandlerFunc,
	roleMiddleware *middleware.RoleMiddleware,
) {
	g := r.Group("exercises")
	{
		g.POST("", jwtMiddleware, roleMiddleware.AllowRole("superadmin", "admin"), exerciseController.Create)
		g.GET("/:exerciseId/score", jwtMiddleware, exerciseController.GetScore)
		g.POST("/:exerciseId/questions", jwtMiddleware, roleMiddleware.AllowRole("superadmin", "admin"), questionController.Create)
		g.POST("/:exerciseId/questions/:questionId/answers", jwtMiddleware, answerController.Create)
	}
}
