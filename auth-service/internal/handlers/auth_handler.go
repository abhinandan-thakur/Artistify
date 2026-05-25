package handlers

import (
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/models"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func Register(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.Users

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err = service.Register(pool, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":    "User Registered",
			"User id":    user.ID,
			"Created at": user.CreatedAt,
		})
	}
}

func RegisterWithRole(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Users

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err = service.RegisterWithRole(pool, user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":    "User Registered",
			"User id":    user.ID,
			"Created at": user.CreatedAt,
			"Acces:":     user.Type,
		})
	}

}

func Login(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.Users

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
			return
		}

		tokenString, role, err := service.Login(pool, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": tokenString, "Acess": role})
	}
}
