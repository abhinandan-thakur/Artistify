package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"context"
	pb "github.com/abhinandan-thakur/Artistify/music-service/proto"
	"log"
)

func AuthMiddleware(authClient pb.AuthenticateServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Entered AuthMiddleware()")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		tokenStrings := strings.TrimPrefix(authHeader, "Bearer ")
		log.Println("Aftere trimPrefix()")
		response, err := authClient.IsValidAndGetRole(context.Background(), &pb.IsValidAndGetRoleRequest{TokenString: tokenStrings})

		if err != nil || !response.Valid {

		c.JSON(http.StatusUnauthorized,	gin.H{"error": "Invalid token",},)
			log.Println("ther error si:", err)
			log.Println("Maybe the response is not vallid", response)
		c.Abort()
		return
		}

		log.Println("Exit Middleware()")
		c.Set("role", response.Role)

		c.Next()



		// claims := jwt.MapClaims{}

		// token, err := jwt.ParseWithClaims(
		// 	tokenStrings,
		// 	claims,
		// 	func(token *jwt.Token) (interface{}, error) {
		// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 			return nil, jwt.ErrTokenSignatureInvalid
		// 		}
		// 		return jwtSecret, nil
		// 	},
		// )

		// if err != nil || !token.Valid {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
		// 	c.Abort()
		// 	return
		// }

		// claims = token.Claims.(jwt.MapClaims)

		// c.Set("id", claims["id"])
		// c.Set("role", claims["role"])
		// c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("entered RequireRole()")
		userRole, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Role doesn't exist"})
			c.Abort()
			return
		}

		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not accessible"})
			c.Abort()
			return
		}
		log.Println("exit rr()")
		c.Next()
	}
}
