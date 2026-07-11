package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/abhinandan-thakur/Artistify/music-service/internal/metrics"
)

func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()

		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}

		method := c.Request.Method
		status := strconv.Itoa(c.Writer.Status())

		// Total requests
		metrics.HttpRequestTotal.
			WithLabelValues(path, status).
			Inc()

		// Error requests
		if c.Writer.Status() >= 400 {
			metrics.HttpRequestErrorTotal.
				WithLabelValues(path, status).
				Inc()
		}

		// Request duration histogram
		metrics.HTTPRequestDuration.
			WithLabelValues(method, path, status).
			Observe(duration)
	}
}