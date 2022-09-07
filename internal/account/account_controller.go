package account

import (
	accountEntities "exercise-api/internal/account/entities"
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
	in := accountEntities.RegisterRequest{}
	authRoleId, _ := c.Get("role_id")
	authUserId, _ := c.Get("user_id")
	in.RequestMetaData.AuthRoleId = int(authRoleId.(float64))
	in.RequestMetaData.AuthUserId = int(authUserId.(float64))

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := a.accountService.Register(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}

func (a *accountController) Login(c *gin.Context) {
	in := accountEntities.LoginRequest{}

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
