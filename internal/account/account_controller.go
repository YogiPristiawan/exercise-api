package account

import (
	"exercise-api/internal/account/entities"
	"exercise-api/internal/presentation"

	"github.com/gin-gonic/gin"
)

type accountController struct {
	stService AccountService
}

func NewAccountController(stService AccountService) *accountController {
	return &accountController{
		stService: stService,
	}
}

func (s *accountController) Register(c *gin.Context) {
	in := entities.RegisterRequest{}

	if !presentation.ReadRestIn(c, &in) {
		return
	}

	out := s.stService.Register(&in)

	presentation.WriteRestOut(c, out, &out.CommonResult)
}
