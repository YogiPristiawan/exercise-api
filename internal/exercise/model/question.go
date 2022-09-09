package model

type GetQuestionById struct {
	Id         int
	ExerciseId int
	Body       string
}

type GetQuestionByExerciseId struct {
}
