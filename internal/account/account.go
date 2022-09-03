package account

import (
	accEntities "exercise-api/internal/account/entities"
	"exercise-api/internal/shared/entities"
)

type AccountService interface {
	Register(*accEntities.RegisterRequest) entities.BaseResponse[accEntities.RegisterResponse]
	Login(*accEntities.LoginRequest) entities.BaseResponse[accEntities.LoginResponse]
	GetAllRole() entities.BaseResponseArray[accEntities.GetAllRoleResponse]
}
