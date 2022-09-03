package routes

import (
	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Register(c *gin.Context)
}

func NewAccountRoutes(r *gin.Engine, controller AccountController) {
	g := r.Group("account")
	{
		g.POST("register", controller.Register)
	}
}
