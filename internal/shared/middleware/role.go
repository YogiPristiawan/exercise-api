package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleMiddleware struct {
	db *gorm.DB
}

func NewRoleMiddleware(db *gorm.DB) *RoleMiddleware {
	return &RoleMiddleware{
		db: db,
	}
}

func (r *RoleMiddleware) AllowRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleId, exists := c.Get("role_id")
		if !exists {
			c.AbortWithStatusJSON(500, gin.H{
				"message": "required jwt middleware",
			})
			return
		}

		// check role in database
		type roleById struct {
			Name string
		}
		role := roleById{}
		err := r.db.Table("roles").Select("name").Where("id = ?", int(roleId.(float64))).First(&role).Error
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "invalid role",
			})
			return
		}

		for _, val := range roles {
			if val == role.Name {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{
			"messagse": "you don't have access to this operation",
		})
	}
}
