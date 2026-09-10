package httpapi

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Prem-Tomar/payment-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(logger *slog.Logger) *gin.Engine {
	router := gin.Default()
	// For Non methods 
	router.HandleMethodNotAllowed = true
	router.NoMethod(methodNotAllow)

	router.Use(middlewares.AddHeaderID)
	router.Use(middlewares.AccessLogger(logger))
	router.Use(requestBodyLimit())

	router.GET("/healthz", healthHandler)
	router.GET("/readyz", readyHandler)

	return router
}

func healthHandler(context *gin.Context) {
	fmt.Println("server started health handler")
	// Retrieve request ID from context
	// requestID, exists := context.Get("request_id")

	// if !exists {
	// 	requestID = "request ID not found"
	// }

	// context.JSON(http.StatusOK, gin.H{
	// 	"status": "ok",
	// })

	writeSuccess(context, http.StatusOK, "ok")
}

func readyHandler(context *gin.Context) {
	fmt.Println("server started")
	// requestID, exists := context.Get("request_id")
	

	// if !exists {
	// 	requestID = "request ID not found"
	// }

	writeSuccess(context, http.StatusOK, "ok")
}

func methodNotAllow(context *gin.Context){

		context.Header("Allow" , "GET")
		writeError(context , http.StatusMethodNotAllowed, "method not allowed")
	}
