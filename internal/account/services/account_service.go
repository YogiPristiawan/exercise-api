package services

import (
	accEntities "exercise-api/internal/account/entities"
	"exercise-api/internal/account/hash"
	"exercise-api/internal/account/jwt"
	"exercise-api/internal/shared/databases"
	"exercise-api/internal/shared/entities"
	"fmt"
)

type accountService struct {
	accountRepository AccountRepository
	validator         Validator
}

var (
	castDatabaseError    = databases.CastDatabaseError
	hashPassword         = hash.HashPassword
	comparePassword      = hash.CompareHashAndPassword
	generateaAccessToken = jwt.GenerateAccessToken
)

func NewAccountService(
	accountRepository AccountRepository,
	validator Validator,
) *accountService {
	return &accountService{
		accountRepository: accountRepository,
		validator:         validator,
	}
}

func (a *accountService) Register(in *accEntities.RegisterRequest) (out entities.BaseResponse[accEntities.RegisterResponse]) {
	// validate payload
	if err := a.validator.ValidateRegisterPayload(in); err != nil {
		out.SetCode(400, err)
		return
	}

	// check if email already exists
	if _, err := a.accountRepository.GetByEmail(in.Email); err != nil {
		if code := castDatabaseError(err); code == 500 {
			out.SetCode(500, err)
			return
		}
	} else {
		out.SetCode(400, fmt.Errorf("email sudah terdaftar"))
		return
	}

	// check available role
	if _, err := a.accountRepository.GetRoleById(in.RoleId); err != nil {
		switch castDatabaseError(err) {
		case 404:
			out.SetCode(400, fmt.Errorf("invalid role_id"))
			return
		case 500:
			out.SetCode(500, err)
			return
		}
	}

	// create user
	hashedPassword, err := hashPassword(in.Password)
	if err != nil {
		out.SetCode(500, err)
		return
	}

	user := accEntities.UserModel{
		Name:     in.Name,
		Email:    in.Email,
		Password: hashedPassword,
		RoleId:   in.RoleId,
	}
	if err := a.accountRepository.Create(&user); err != nil {
		out.SetCode(500, err)
		return
	}

	out.Message = "berhasil mendaftar silahkan login!"
	out.SetCode(201, nil)

	return
}

func (a *accountService) Login(in *accEntities.LoginRequest) (out entities.BaseResponse[accEntities.LoginResponse]) {
	// validate request
	if err := a.validator.ValidateLoginPayload(in); err != nil {
		out.SetCode(400, err)
		return
	}

	// check if email exists
	user, err := a.accountRepository.GetByEmail(in.Email)
	switch castDatabaseError(err) {
	case 404:
		out.SetCode(404, fmt.Errorf("email belum terdaftar"))
		return
	case 500:
		out.SetCode(500, err)
		return
	}

	// verify password
	if err := comparePassword(user.Password, in.Password); err != nil {
		out.SetCode(400, fmt.Errorf("password yang anda masukkan salah"))
		return
	}

	// generate token
	accessToken, err := generateaAccessToken(&accEntities.JwtClaims{
		UserId: user.Id,
		RoleId: user.RoleId,
	})
	if err != nil {
		out.SetCode(500, err)
		return
	}

	out.Data.AccessToken = accessToken
	return
}

func (a *accountService) GetAllRole() (out entities.BaseResponseArray[accEntities.GetAllRoleResponse]) {
	// get all role
	roles, err := a.accountRepository.GetAllRole()
	switch castDatabaseError(err) {
	case 500:
		out.SetCode(500, err)
		return
	}

	out.Message = "list roles"
	if len(roles) > 0 {
		for _, val := range roles {
			out.Data = append(out.Data, &accEntities.GetAllRoleResponse{
				Id:   val.Id,
				Name: val.Name,
			})
		}
	} else {
		out.Data = []*accEntities.GetAllRoleResponse{}
	}
	return
}
