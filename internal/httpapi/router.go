package httpapi

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Prem-Tomar/payment-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(logger *slog.Logger) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery()) // added as it was in default not it new , for panic handling
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
	writeSuccess(context, http.StatusOK, "ok")
}

func readyHandler(context *gin.Context) {
	fmt.Println("server started")
	writeSuccess(context, http.StatusOK, "ready")
}

func methodNotAllow(context *gin.Context) {

	context.Header("Allow", "GET")
	writeError(context, http.StatusMethodNotAllowed, "method not allowed")
}
