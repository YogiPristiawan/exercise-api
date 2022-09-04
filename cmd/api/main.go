package main

import (
	"exercise-api/internal/account"
	accountRepository "exercise-api/internal/account/repositories"
	accountService "exercise-api/internal/account/services"
	accountValidator "exercise-api/internal/account/validator"
	"exercise-api/internal/exercise"
	exerciseRepository "exercise-api/internal/exercise/repositories"
	exerciseService "exercise-api/internal/exercise/services"
	exerciseValidator "exercise-api/internal/exercise/validator"
	"exercise-api/internal/routes"
	"exercise-api/internal/shared/databases"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	// initialize databases
	mysql := databases.NewMySQLConn()

	// initialize repositories
	userRepositoryImpl := accountRepository.NewUserRepository(mysql)
	exerciseRepositoryImpl := exerciseRepository.NewExerciseRepository(mysql)
	questionRepositoryImpl := exerciseRepository.NewQuestionRepository(mysql)

	// initialize validators
	validator := validator.New()
	accountValidatorImpl := accountValidator.NewAccountValidator(validator)
	exerciseValidatorImpl := exerciseValidator.NewExerciseValidator(validator)
	questionValidatorImpl := exerciseValidator.NewQuestionValidator(validator)

	// initialize services
	accountServiceImpl := accountService.NewAccountService(userRepositoryImpl, accountValidatorImpl)
	exerciseServiceImpl := exerciseService.NewExerciseService(exerciseValidatorImpl, exerciseRepositoryImpl)
	questionServiceImpl := exerciseService.NewQuestionService(questionValidatorImpl, questionRepositoryImpl, exerciseRepositoryImpl)

	// initialize controllers
	accountController := account.NewAccountController(accountServiceImpl)
	exerciseController := exercise.NewExerciseController(exerciseServiceImpl)
	questionController := exercise.NewQuestionController(questionServiceImpl)

	// initilaize routes
	routes.NewAccountRoutes(router, accountController)
	routes.NewExerciseRoutes(router, exerciseController, questionController)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
