package main

import (
	"artistify/internal/database"
	"artistify/internal/handlers"
	"artistify/internal/middleware"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"rsc.io/quote"
	"strconv"
)

func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"quote": quote.Go(),
	})
}

// Define metrics
var (
	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "Total number of requests processed by the API",
	}, []string{"path", "status"})

	HttpRequestErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_error_total",
		Help: "Total number of errors returned by the API",
	}, []string{"path", "status"})
)

func init() {
	prometheus.MustRegister(
		HttpRequestTotal,
		HttpRequestErrorTotal,
	)
}

// Middleware to record incoming requests metrics
func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.Next()
		status := c.Writer.Status()
		if status < 400 {
			HttpRequestTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		} else {
			HttpRequestErrorTotal.WithLabelValues(path, strconv.Itoa(status)).Inc()
		}
	}
}

var conn *pgx.Conn

func main() {

	var err error

	conn, err = database.ConnectDB()
	redisClient := database.ConnectRedis()

	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())

	fmt.Println("DATABASE Successfully connected!!!")

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	fmt.Println("Running server at 8080...", quote.Go())

	ipRateLimiter := middleware.NewIPRateLimiter()
	router.Use(RequestMetricsMiddleware())

	router.GET("/", home)
	//  router.GET("/metrics", PrometheusHandler())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/auth/login", middleware.RateLimitMiddleware(ipRateLimiter), handlers.Login(conn))
	router.POST("/auth/register", handlers.Register(conn))
	router.POST("/auth/registerWithRole", handlers.RegisterWithRole(conn))
	router.GET("/albums", handlers.GetAlbums(conn, redisClient))
	router.GET("/albums/artist/:artist", handlers.GetAlbumsByArtist(conn))
	router.GET("/albums/:id", handlers.GetAlbumByID(conn))

	artist := router.Group("/artist")
	artist.Use(middleware.AuthMiddleware())
	artist.Use(middleware.RequireRole("artist"))
	{
		artist.POST("/albums", handlers.PostAlbum(conn))
		artist.PUT("/albums/:id", handlers.UpdateAlbumByID(conn))
		artist.DELETE("/albums/:id", handlers.DeleteAlbumByID(conn))
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequireRole("admin"))
	{
		admin.POST("/albums", handlers.PostAlbum(conn))
		admin.PUT("/albums/:id", handlers.UpdateAlbumByID(conn))
		admin.DELETE("/albums/:id", handlers.DeleteAlbumByID(conn))
	}

	// artist.GET("/albums/stats/:id", albumStatsByID)

	router.Run(":8080")
}
