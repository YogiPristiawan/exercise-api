package entities

import "exercise-api/internal/shared/entities"

// register response
type RegisterRequest struct {
	entities.RequestMetaData
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id" validate:"required"`
}

type RegisterResponse struct {
	Id int `json:"id"`
}

// login response
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

// get all role response
type GetAllRoleResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
