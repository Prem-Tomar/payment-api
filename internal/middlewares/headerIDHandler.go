package middlewares

import (

	"github.com/gin-gonic/gin"
	 "github.com/google/uuid"
)

func AddHeaderID(c *gin.Context) {
requestID := c.GetHeader("X-Request-ID")

		// Generate a new one if the header is missing
		if requestID == "" {
			requestID = uuid.New().String() 
		
		}

		// Store in context for handlers and echo back in response headers
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
}