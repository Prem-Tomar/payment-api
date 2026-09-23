package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireAPIKey(c *gin.Context) {

	apiKey := strings.TrimSpace(c.GetHeader("X-API-Key"))

	if apiKey == "" {

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "api key is required"})
		return
	}
	c.Set("api_key", apiKey)
	c.Next()
}
