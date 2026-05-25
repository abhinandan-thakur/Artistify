package main

import (
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/database"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/handlers"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/middleware"
	"fmt"
	"log"
	"os"
	grpcserver "github.com/abhinandan-thakur/Artistify/auth-service/internal/grpc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"

)

var pool *pgxpool.Pool

func main() {

	var err error

	config := config.LoadConfig()

	pool, err = database.ConnectDB(config)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer pool.Close()

	fmt.Println("DATABASE Successfully connected!!!")

	err = grpcserver.StartGRPCServer(config)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	router := gin.Default()

	loginRateLimiter := middleware.NewIPRateLimiter(5.00, 0.50)
	registerRateLimiter := middleware.NewIPRateLimiter(3.00, 0.33)
	adminRateLimiter := middleware.NewIPRateLimiter(2.00, 0.05)
	router.Use(middleware.RequestMetricsMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/auth/login", middleware.RateLimitMiddleware(loginRateLimiter), handlers.Login(pool))
	router.POST("/auth/register", middleware.RateLimitMiddleware(registerRateLimiter), handlers.Register(pool))
	router.POST("/auth/registerWithRole", middleware.RateLimitMiddleware(adminRateLimiter), handlers.RegisterWithRole(pool))

	port := ":"+os.Getenv("URL_PORT")
	err = router.Run(port)
	log.Println("Auth Service running at "+os.Getenv("URL_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}

