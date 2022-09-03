package entities

// register response
type RegisterRequest struct {
	Name     string `json:"name" validate:"required,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required"`
	RoleId   int    `json:"role_id" validate:"required"`
}

type RegisterResponse struct {
	Message string `json:"message"`
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
