package httpapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/healthz", readyHandler)
	router.GET("/readyz", healthHandler)

	return router
}

func healthHandler(context *gin.Context) {
	fmt.Println("server started")

	context.JSON(http.StatusOK, gin.H{
		"Statuf_Gin_Server": "Server Started",
	})
}

func readyHandler(context *gin.Context) {
	fmt.Println("server started")

	context.JSON(http.StatusOK, gin.H{
		"Statuf_Gin_Server": "Server Started",
	})
}