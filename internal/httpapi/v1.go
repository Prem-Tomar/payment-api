package httpapi

import( "github.com/gin-gonic/gin"
"fmt")

func registerV1Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
fmt.Println(v1)
	// Future API ke liye
}
