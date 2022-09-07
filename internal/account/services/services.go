package services

import (
	"exercise-api/internal/account/entities"
	"exercise-api/internal/account/model"
)

type AccountRepository interface {
	GetByEmail(string) (*model.GetByEmail, error)
	GetRoleById(int) (*model.GetRoleById, error)
	Create(*entities.UserModel) error
	GetAllRole() ([]*model.GetAllRole, error)
}

type Validator interface {
	ValidateRegisterPayload(*entities.RegisterRequest) error
	ValidateLoginPayload(*entities.LoginRequest) error
}
