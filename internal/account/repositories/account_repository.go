package repositories

import (
	"exercise-api/internal/account/entities"
	"exercise-api/internal/account/model"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *accountRepository {
	return &accountRepository{
		db: db,
	}
}

func (u *accountRepository) GetByEmail(email string) (user *model.GetByEmail, err error) {
	err = u.db.Table("users").Select(
		"id",
		"email",
		"password",
		"role_id",
	).Where("email = ?", email).First(&user).Error
	return
}

func (u *accountRepository) GetRoleById(id int) (role *model.GetRoleById, err error) {
	err = u.db.Table("roles").Select("id", "name").Where("id = ?", id).First(&role).Error
	return
}

func (u *accountRepository) Create(user *entities.UserModel) (err error) {
	err = u.db.Create(user).Error
	return
}

func (u *accountRepository) GetAllRole() (roles []*model.GetAllRole, err error) {
	err = u.db.Table("roles").First(&roles).Error
	return
}
