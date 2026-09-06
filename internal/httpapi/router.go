package httpapi

import (
	"fmt"
	"net/http"

	"github.com/Prem-Tomar/payment-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	  router.Use(middlewares.AddHeaderID)
	router.GET("/healthz", healthHandler)
     router.GET("/readyz", readyHandler)

	return router
}

func healthHandler(context *gin.Context) {
	fmt.Println("server started health handler")
// Retrieve request ID from context
		 requestID, exists := context.Get("request_id")

		  if !exists {
        requestID = "request ID not found"
    }

	context.JSON(http.StatusOK, gin.H{
		"Statuf_Gin_Server": "Server Started",
		"request_Id" : requestID,
	})
}

func readyHandler(context *gin.Context) {
	fmt.Println("server started")
 requestID, exists := context.Get("request_id")

		  if !exists {
        requestID = "request ID not found"
    }

	context.JSON(http.StatusOK, gin.H{
		"Statuf_Gin_Server": "Server Started",
			"request_Id" : requestID,
	})
}