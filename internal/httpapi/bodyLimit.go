package httpapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const maxRequestBodySize = 1 << 20 // This is size less than 1 MB

func requestBodyLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > maxRequestBodySize {
			//Reject immediatly as content length > max required size
			writeError(
				c,
				http.StatusRequestEntityTooLarge,
				"request body too large",
			)

			// to not allow it to go another middlewares
			c.Abort()
			return
		}
		// Got the Idea that , what if data is coming in chunks and is more that the required size ,  we will handle it in POST request when created
		c.Request.Body = http.MaxBytesReader(
			c.Writer,
			c.Request.Body,
			maxRequestBodySize,
		)
		c.Next()
	}
}
