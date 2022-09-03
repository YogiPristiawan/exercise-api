package routes

import (
	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Role(c *gin.Context)
}

func NewAccountRoutes(r *gin.Engine, controller AccountController) {
	g := r.Group("account")
	{
		g.GET("roles", controller.Role)
		g.POST("register", controller.Register)
		g.POST("login", controller.Login)
	}
}
