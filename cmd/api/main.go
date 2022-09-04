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
	userRepository := accountRepository.NewUserRepository(mysql)
	exerciseRepository := exerciseRepository.NewExerciseRepository(mysql)

	// initialize validators
	validator := validator.New()
	accountValidator := accountValidator.NewAccountValidator(validator)
	exerciseValidator := exerciseValidator.NewExerciseValidator(validator)

	// initialize services
	accountService := accountService.NewAccountService(userRepository, accountValidator)
	exerciseService := exerciseService.NewExerciseService(exerciseValidator, exerciseRepository)

	// initialize controllers
	accountController := account.NewAccountController(accountService)
	exerciseController := exercise.NewExerciseController(exerciseService)

	// initilaize routes
	routes.NewAccountRoutes(router, accountController)
	routes.NewExerciseRoutes(router, exerciseController)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
