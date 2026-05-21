package middleware

import (
	"strings"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"fmt"
)

var jwtSecret = []byte("super-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		tokenStrings := strings.TrimPrefix(authHeader, "Bearer ")
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(
			tokenStrings, 
			claims, 
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); ! ok {
					return nil, jwt.ErrTokenSignatureInvalid
				}
			return jwtSecret, nil
		},
	)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		claims = token.Claims.(jwt.MapClaims)
		
		c.Set("id", claims["id"])
        c.Set("role", claims["role"])
		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Role doesn't exist"})
			c.Abort()
			return
		}

		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"error":"Not accessible"})
			c.Abort()
			return
		}

		c.Next()
	}
}
