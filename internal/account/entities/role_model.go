package entities

type RoleModel struct {
	Id   int `gorm:"primaryKey"`
	Name string
}
