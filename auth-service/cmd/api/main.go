// @title Authentication Service
// @version 1.0
// @description authentication service for Artistify app
// @host localhost:8180
package main

import (
	"fmt"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/database"
	grpc "github.com/abhinandan-thakur/Artistify/auth-service/internal/grpc"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/handlers"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/abhinandan-thakur/Artistify/auth-service/docs"
)

var pool *pgxpool.Pool

func main() {

	var err error

	config := config.LoadConfig()
	fmt.Println("The environment is:", config.AppEnv)
	fmt.Println("The JWT secret is:", config.JWTSecret)
	fmt.Println("The rate_limit_login is:", config.LoginRateLimit, "==", config.LoginRateRefil)
	fmt.Println("The reigster is:", config.RegisterRateLimit, "==", config.ReigsterRateRefil)
	fmt.Println("The admin is:", config.AdminRateLimit, "==", config.AdminRateRefil)

	pool, err = database.ConnectDB(config)
	redisClient := database.ConnectRedis(config)
	fmt.Println("Redis connected...")

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer pool.Close()

	fmt.Println("DATABASE Successfully connected!!!")

	err = grpc.StartGRPCServer(config)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println("Started GRPC auth server")

	mailingClient, err := grpc.NewMailingClient()

	if err != nil {
		log.Fatal(
			"failed to connect mailing service:",
			err,
		)
	}

	fmt.Println("Started mailing client")

	router := gin.Default()

	loginRateLimiter := middleware.NewIPRateLimiter(config.LoginRateLimit, config.LoginRateRefil)
	registerRateLimiter := middleware.NewIPRateLimiter(config.RegisterRateLimit, config.ReigsterRateRefil)
	adminRateLimiter := middleware.NewIPRateLimiter(config.AdminRateLimit, config.AdminRateRefil)
	router.Use(middleware.RequestMetricsMiddleware())

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/auth/login", middleware.RateLimitMiddleware(loginRateLimiter), handlers.Login(pool))
	router.POST("/auth/sendOTP", middleware.RateLimitMiddleware(registerRateLimiter), handlers.SendOTP(pool, redisClient, mailingClient))
	router.POST("/auth/register", middleware.RateLimitMiddleware(registerRateLimiter), handlers.Register(pool))
	router.POST("/auth/registerWithRole", middleware.RateLimitMiddleware(adminRateLimiter), handlers.RegisterWithRole(pool))

	port := ":" + os.Getenv("URL_PORT")
	err = router.Run(port)
	log.Println("Auth Service running at " + os.Getenv("URL_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
