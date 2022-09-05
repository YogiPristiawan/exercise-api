package routes

import "github.com/gin-gonic/gin"

type ExerciseController interface {
	Create(c *gin.Context)
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
) {
	g := r.Group("exercises")
	{
		g.POST("", exerciseController.Create)
		g.POST("/:exerciseId/questions", questionController.Create)
		g.POST("/:exerciseId/questions/:questionId/answers", answerController.Create)
	}
}
