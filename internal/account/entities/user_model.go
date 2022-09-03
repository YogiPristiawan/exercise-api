package entities

type UserModel struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	RoleId    int
	CreatedAt int `gorm:"autoCreateTime"`
	UpdatedAt int `gorm:"autoUpdateTime"`
}

func (u *UserModel) TableName() string {
	return "users"
}
