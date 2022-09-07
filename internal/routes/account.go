package routes

import (
	"exercise-api/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Role(c *gin.Context)
}

func NewAccountRoutes(
	r *gin.Engine,
	controller AccountController,
	jwtMiddleware gin.HandlerFunc,
	roleMiddleware *middleware.RoleMiddleware,
) {
	g := r.Group("account")
	{
		g.GET("roles", controller.Role)
		g.POST("register", jwtMiddleware, roleMiddleware.AllowRole("superadmin", "admin"), controller.Register)
		g.POST("login", controller.Login)
	}
}
