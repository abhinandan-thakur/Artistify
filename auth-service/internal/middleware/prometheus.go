package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/abhinandan-thakur/Artistify/auth-service/internal/metrics"
)

func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
		path := c.FullPath()

		status := strconv.Itoa(c.Writer.Status())

		if c.Writer.Status() < 400 {
			metrics.HttpRequestTotal.WithLabelValues(path, status).Inc()
		} else {
			metrics.HttpRequestErrorTotal.WithLabelValues(path, status,).Inc()
		}
	}
}