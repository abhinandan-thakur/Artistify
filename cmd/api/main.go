package main

import (
	"artistify/internal/database"
	"artistify/internal/handlers"
	"artistify/internal/middleware"
	// "context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
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

var pool *pgxpool.Pool

func main() {

	var err error

	pool, err = database.ConnectDB()
	defer pool.Close()
	redisClient := database.ConnectRedis()

	if err != nil {
		panic(err)
	}

	fmt.Println("DATABASE Successfully connected!!!")

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	fmt.Println("Running server at 8080...", quote.Go())

	ipRateLimiter := middleware.NewIPRateLimiter()
	router.Use(RequestMetricsMiddleware())

	router.GET("/", home)
	//  router.GET("/metrics", PrometheusHandler())
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/auth/login", middleware.RateLimitMiddleware(ipRateLimiter), handlers.Login(pool))
	router.POST("/auth/register", handlers.Register(pool))
	router.POST("/auth/registerWithRole", handlers.RegisterWithRole(pool))
	router.GET("/albums", handlers.GetAlbums(pool, redisClient))
	router.GET("/albums/artist/:artist", handlers.GetAlbumsByArtist(pool))
	router.GET("/albums/:id", handlers.GetAlbumByID(pool))

	artist := router.Group("/artist")
	artist.Use(middleware.AuthMiddleware())
	artist.Use(middleware.RequireRole("artist"))
	{
		artist.POST("/albums", handlers.PostAlbum(pool))
		artist.PUT("/albums/:id", handlers.UpdateAlbumByID(pool))
		artist.DELETE("/albums/:id", handlers.DeleteAlbumByID(pool))
	}

	admin := router.Group("/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RequireRole("admin"))
	{
		admin.POST("/albums", handlers.PostAlbum(pool))
		admin.PUT("/albums/:id", handlers.UpdateAlbumByID(pool))
		admin.DELETE("/albums/:id", handlers.DeleteAlbumByID(pool))
	}

	// artist.GET("/albums/stats/:id", albumStatsByID)

	router.Run(":8080")
}
