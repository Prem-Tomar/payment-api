package middlewares

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recovery(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil { // recover captures panic
				logger.Error(
					"panic recovered",
					"panic", recovered,
					"stack", string(debug.Stack()),
				)
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{
						"error": "internal server error",
					},
				)
			}

		}()
		c.Next()
	}
}
