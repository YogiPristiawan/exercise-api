package account

import (
	"exercise-api/internal/account/entities"
	"exercise-api/internal/presentation"

	"github.com/gin-gonic/gin"
)

type accountController struct {
	accountService AccountService
}

func NewAccountController(accountService AccountService) *accountController {
	return &accountController{
		accountService: accountService,
	}
}

func (a *accountController) Register(c *gin.Context) {
	in := entities.RegisterRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := a.accountService.Register(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (a *accountController) Login(c *gin.Context) {
	in := entities.LoginRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := a.accountService.Login(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (a *accountController) Role(c *gin.Context) {
	out := a.accountService.GetAllRole()

	presentation.WriteRestOut(c, out, &out.CommonResult)
}
