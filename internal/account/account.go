package account

import (
	"exercise-api/internal/account/entities"
)

type AccountService interface {
	Register(*entities.RegisterRequest) entities.RegisterResponse
}
