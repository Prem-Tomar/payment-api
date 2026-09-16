package httpapi

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Prem-Tomar/payment-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func newRouter(logger *slog.Logger, checker ReadinessChecker) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery()) // added as it was in default not it new , for panic handling
	// For Non methods
	router.HandleMethodNotAllowed = true
	router.NoMethod(methodNotAllow)

	// Middlewares
	router.Use(middlewares.AddHeaderID)
	router.Use(middlewares.AccessLogger(logger))
	router.Use(requestBodyLimit())

	// Groups
	registerV1Routes(router)
	// future public APIs

	// checker := DefaultReadinessChecker{}

	router.GET("/healthz", healthHandler)
	router.GET("/readyz", readyHandler(checker))

	router.NoRoute(noRouteHandler)

	return router
}

func NewRouter(logger *slog.Logger) *gin.Engine {
	checker := DefaultReadinessChecker{}

	return newRouter(logger, checker)
}

func healthHandler(context *gin.Context) {
	fmt.Println("server started health handler")
	writeSuccess(context, http.StatusOK, "ok")
}

func readyHandler(checker ReadinessChecker) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := checker.Check(); err != nil {
			writeError(context, http.StatusServiceUnavailable, "service unavailable")
			return
		}
		writeSuccess(context, http.StatusOK, "ready")
	}
}

func methodNotAllow(context *gin.Context) {

	context.Header("Allow", "GET")
	writeError(context, http.StatusMethodNotAllowed, "method not allowed")
}

func noRouteHandler(context *gin.Context) {
	writeError(context, http.StatusNotFound, "route not found")
}
