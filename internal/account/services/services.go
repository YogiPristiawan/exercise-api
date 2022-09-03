package services

import (
	"exercise-api/internal/account/entities"
	"exercise-api/internal/account/shared"
)

type AccountRepository interface {
	GetByEmail(string) (*shared.GetByEmailDTO, error)
	GetRoleById(int) (*shared.GetRoleByIdDTO, error)
	Create(*entities.UserModel) error
	GetAllRole() ([]*shared.GetAllRoleDTO, error)
}

type Validator interface {
	ValidateRegisterPayload(*entities.RegisterRequest) error
	ValidateLoginPayload(*entities.LoginRequest) error
}
