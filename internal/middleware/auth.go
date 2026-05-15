package middleware


import (
	"strings"
	"net/http"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
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


/*
func RBACMiddleware(permission string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, exists := c.Get("id")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
            c.Abort()
            return
        }
if !database.HasAccess(userID.(int), permission) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
            c.Abort()
            return
        }
        c.Next()
    }
}

func HasAccess(userID int, accessName string) bool {
    var count int
    query := `SELECT COUNT(*)
    FROM users u
    JOIN user_role ur ON u.id = ur.user_id
    JOIN role_access ra ON ur.role_id = ra.role_id
    JOIN access a ON ra.access_id = a.access_id
    WHERE u.id = $1 AND a.access_name = $2`

    err := Db.QueryRow(query, userID, accessName).Scan(&count)
    if err != nil {
        // Handle error (optional: log the error)
        return false
    }

    return count > 0
}

*/