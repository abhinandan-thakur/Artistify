package main

import (
	grpcclient "github.com/abhinandan-thakur/Artistify/music-service/internal/grpc"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/database"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/handlers"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/middleware"
	"github.com/abhinandan-thakur/Artistify/music-service/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"os"
)

var pool *pgxpool.Pool

func main() {

	config := config.LoadConfig()

	var err error

	pool, err = database.ConnectDB(config)
	redisClient := database.ConnectRedis(config)

	authClient := grpcclient.NewAuthClient()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer pool.Close()

	fmt.Println("DATABASE Successfully connected!!!")

	router := gin.Default()

	router.Use(middleware.RequestMetricsMiddleware())

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	getRateLimiter    := middleware.NewIPHash(50.00, 3.00)
	postRateLimiter	  := middleware.NewIPHash(3.00, 0.50)
	deleteRateLimiter := middleware.NewIPHash(1.00, 0.20)
	putRateLimiter	  := middleware.NewIPHash(4.00, 0.60)

	router.GET("/albums", middleware.RateLimitingMiddleware(getRateLimiter), handlers.GetAlbums(pool, redisClient))
	router.GET("/albums/artist/:artist", middleware.RateLimitingMiddleware(getRateLimiter), handlers.GetAlbumsByArtist(pool))
	router.GET("/albums/:id", middleware.RateLimitingMiddleware(getRateLimiter), handlers.GetAlbumByID(pool))

	artist := router.Group("/artist")
	artist.Use(middleware.AuthMiddleware(authClient))
	artist.Use(middleware.RequireRole("artist"))
	{
		artist.POST("/albums", middleware.RateLimitingMiddleware(postRateLimiter) ,handlers.PostAlbum(pool))
		artist.PUT("/albums/:id", middleware.RateLimitingMiddleware(putRateLimiter), handlers.UpdateAlbumByID(pool))
		artist.DELETE("/albums/:id", middleware.RateLimitingMiddleware(deleteRateLimiter), handlers.DeleteAlbumByID(pool))
	}

	admin := router.Group("/admin")

	admin.Use(middleware.AuthMiddleware(authClient,))
	admin.Use(middleware.RequireRole("admin"))
	{
		admin.POST("/albums", middleware.RateLimitingMiddleware(postRateLimiter), handlers.PostAlbum(pool))
		admin.PUT("/albums/:id", middleware.RateLimitingMiddleware(putRateLimiter),handlers.UpdateAlbumByID(pool))
		admin.DELETE("/albums/:id", middleware.RateLimitingMiddleware(deleteRateLimiter),handlers.DeleteAlbumByID(pool))
	}

	port := ":"+os.Getenv("URL_PORT")
	err = router.Run(port)
	log.Println("Auth Service running at "+os.Getenv("URL_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}