package httpapi

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const maxRequestBodySize = 1 << 20 // This is size less than 1 MB

func requestBodyLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		const maxBodySize = 1 << 20

		body, err := io.ReadAll(
			io.LimitReader(c.Request.Body, maxBodySize+1),
		)

		if err != nil {
			writeError(c, http.StatusBadRequest, "unable to read request body")
			c.Abort()
			return
		}

		if len(body) > maxBodySize {
			writeError(c, http.StatusRequestEntityTooLarge, "request body too large")
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		c.Next()
	}
}
