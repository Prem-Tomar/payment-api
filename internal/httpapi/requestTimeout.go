package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const publicRequestTimeOut = 5 * time.Second

func requestTimeOut(timeout time.Duration) gin.HandlerFunc {
return func(c *gin.Context){

	ctx , cancel := context.WithTimeout(c.Request.Context() , timeout,)

	defer cancel()

	c.Request = c.Request.WithContext(ctx)

	c.Next()

	if ctx.Err() == context.DeadlineExceeded && c.Writer.Written() { 
		writeError(
			c,
			http.StatusGatewayTimeout,
			"request timed out",
			)
		}
}
}