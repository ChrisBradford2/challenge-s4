package middleware

import (
	"challenges4/config"
	"challenges4/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func AuthMiddleware(requiredRole uint8) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := jwt.ParseWithClaims(tokenString, &services.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(*services.Claims); ok && token.Valid {
			if claims.Roles&requiredRole != requiredRole {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
				return
			}
			c.Set("userID", claims.UserID)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
	}
}
