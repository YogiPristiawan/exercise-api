package entities

import "exercise-api/internal/shared/entities"

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id" validate:"required"`
}

type RegisterResponse struct {
	entities.CommonResult
	Message string `json:"message"`
}
