package handlers

import (
	grpc "github.com/abhinandan-thakur/Artistify/auth-service/internal/grpc"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/models"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type RegisterResponse struct {
	Message   string    `json:"message"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterWithRoleResponse struct {
	Message   string    `json:"message"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	Role      string    `json:"role"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type SendOTPResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

// Register godoc
//
// @Summary Register
// @Description Registers the user into the database
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Param request body models.Users true "Register User"
//
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /auth/register [post]
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

// RegisterWithRole godoc
//
// @Summary Register the user and also let them select the role
// @Description Registers the user also accepts role as a json field
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Param request body models.Users true "RegisterWithRole User"
//
// @Success 201 {object} RegisterWithRoleResponse
// @Failure 400 {object} ErrorResponse
//
// @Router /auth/registerWithRole [post]
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
			"Role":       user.Type,
		})
	}

}

// Login godoc
//
// @Summary Login
// @Description Authenticate a user and return JWT tokens.
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Param request body models.Users true "Login User"
//
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /auth/login [post]
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

		c.JSON(http.StatusOK, gin.H{"token": tokenString, "Acess": role})
	}
}

// SendOTP godoc
//
// @Summary Send otp
// @Description hashes a secure random otp and sends a mail to the json input field
// @Tags Authentication
// @Accept json
// @Produce json
//
// @Param request body models.Users true "Send otp Request"
//
// @Success 202 {object} SendOTPResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
//
// @Router /auth/sendOTP [post]
func SendOTP(pool *pgxpool.Pool, rdb *redis.Client, mailingClient *grpc.MailingClient) gin.HandlerFunc {

	return func(c *gin.Context) {
		var user models.Users

		err := c.BindJSON(&user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		otp, err := service.GenerateOTP(user, rdb)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = mailingClient.SendMail(user.Email, otp)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"message": "otp queued"})
	}
}
