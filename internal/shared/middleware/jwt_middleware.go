package middleware

import (
	"exercise-api/internal/shared/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	decodeJwt = jwt.DecodeJwt
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer") {
			c.AbortWithStatusJSON(401, map[string]string{
				"message": "unauthorize",
			})
			return
		}

		token := strings.Split(header, " ")[1]

		claims, err := decodeJwt(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.Set("user_id", claims["user_id"])
		c.Set("role_id", claims["role_id"])
		c.Next()
	}
}
