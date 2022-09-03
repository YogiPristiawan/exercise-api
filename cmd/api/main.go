package main

import (
	"exercise-api/internal/account"
	accountRepository "exercise-api/internal/account/repositories"
	accountService "exercise-api/internal/account/services"
	accountValidator "exercise-api/internal/account/validator"
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

	// initialize validators
	validator := validator.New()
	accountValidator := accountValidator.NewAccountValidator(validator)

	// initialize services
	accountService := accountService.NewAccountService(userRepository, accountValidator)

	// initialize controllers
	accountController := account.NewAccountController(accountService)

	// initilaize routes
	routes.NewAccountRoutes(router, accountController)

	router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
