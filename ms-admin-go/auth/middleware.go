package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware extrae claims y los guarda
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		claims, err := ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Set("userId", claims.UserID.String())
		c.Set("role", claims.Role)
		c.Next()
	}
}
