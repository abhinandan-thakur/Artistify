package handlers

import (
	"github.com/jackc/pgx/v5"
	"artistify/internal/models"
	"artistify/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {

		
		var user models.Users
		
		err := c.BindJSON(&user)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err = service.Register(conn, user)
		
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	

		c.JSON(http.StatusCreated, gin.H{
				"message":"User Registered",
				"User id":user.ID,
				"Created at":user.CreatedAt,
		})
	}
}

func RegisterWithRole(conn *pgx.Conn) gin.HandlerFunc { 
	return func(c *gin.Context)	{
		var user models.Users
		
		err := c.BindJSON(&user)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err = service.RegisterWithRole(conn, user)
		
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		c.JSON(http.StatusCreated, gin.H{
			"message":"User Registered",
			"User id": user.ID,
			"Created at": user.CreatedAt,
			"Acces:": user.Type,
		})
	}
	
}

func Login(conn *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
			
		var user models.Users
			
		err := c.BindJSON(&user)
			
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid Input"})
			return
		}

		tokenString, role, err := service.Login(conn, user)
			
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"could not create token"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"token": tokenString, "Acess": role})
	}
}

