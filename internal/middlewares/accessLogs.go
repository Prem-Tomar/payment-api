package middlewares

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func AccessLogger(logger *slog.Logger) gin.HandlerFunc {

	return func(c *gin.Context) {
		start := time.Now()
		// Intentional Gap as this will allow other Middlewares to execute
		c.Next()

		requestID, _ := c.Get("request_id")

		logger.Info("HTTP request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"duration", time.Since(start),
			"request_id", requestID,
		)
	}
}
