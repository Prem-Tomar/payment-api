package httpapi

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const maxRequestBodySize = 1 <<20 // This is size less than 1 MB

func requestBodyLimit() gin.HandlerFunc {
	return  func(c *gin.Context) {
		if c.Request.ContentLength >  maxRequestBodySize {
			writeError(
				c ,
				http.StatusRequestEntityTooLarge,
				"request body too large",
			)

			// to not allow it to go another middlewares
			c.Abort()
			return
		}

		c.Request.Body = http.MaxBytesReader(
			c.Writer,
			c.Request.Body,
			maxRequestBodySize,
		)
		c.Next()
	}
}