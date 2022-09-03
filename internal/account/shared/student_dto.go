package shared

type GetByEmailDTO struct {
	Id       int
	Email    string
	Password string
	RoleId   int
}
