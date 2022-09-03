package services

import (
	"exercise-api/internal/account/entities"
	"exercise-api/internal/shared/databases"
	"exercise-api/internal/shared/hash"
	"fmt"
)

type accountService struct {
	accountRepository AccountRepository
	validator         Validator
}

var (
	castDatabaseError = databases.CastDatabaseError
	hashPassword      = hash.HashPassword
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

func (s *accountService) Register(in *entities.RegisterRequest) (out entities.RegisterResponse) {
	// validate payload
	if err := s.validator.ValidateRegisterPayload(in); err != nil {
		out.SetCode(400, err)
		return
	}

	// check if email already exists
	if _, err := s.accountRepository.GetByEmail(in.Email); err != nil {
		if code := castDatabaseError(err); code == 500 {
			out.SetCode(500, err)
			return
		}
	} else {
		out.SetCode(400, fmt.Errorf("email sudah terdaftar"))
		return
	}

	// check available role
	if _, err := s.accountRepository.GetRoleById(in.RoleId); err != nil {
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

	user := entities.UserModel{
		Name:     in.Name,
		Email:    in.Email,
		Password: hashedPassword,
		RoleId:   in.RoleId,
	}
	if err := s.accountRepository.Create(&user); err != nil {
		out.SetCode(500, err)
		return
	}

	out.Message = "berhasil mendaftar silahkan login!"
	out.SetCode(201, nil)

	return
}
