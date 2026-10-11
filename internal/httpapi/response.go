package httpapi

import (
	"github.com/gin-gonic/gin"
)

type APIErrorCode string

const (
	ErrInvalidRequest  APIErrorCode = "INVALID_REQUEST"
	ErrRequestTooLarge APIErrorCode = "REQUEST_TOO_LARGE"
	ErrValidation      APIErrorCode = "VALIDATION_ERR"
	ErrNotImplemented  APIErrorCode = "NOT_IMPLEMENTED"
	ErrInternal        APIErrorCode = "INTERNAL_ERROR"
	ErronNotFound      APIErrorCode = "NOT_FOUND"
	ErrGatewayTimeout APIErrorCode = "GATEWAY_TIMEOUT"
	ErrMethodNotAllowed APIErrorCode = "METHOD_NOT_ALLOWED"
	ErrServiceUnavailable APIErrorCode = "SERVICE_UNAVAILABLE"
	
)

type successResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Code  APIErrorCode `json:"code"`
	Error string       `json:"error"`
}

func writeSuccess(c *gin.Context, status int, message string) {
	c.JSON(status, successResponse{
		Status: message,
	})
}

func writeError(c *gin.Context, status int, code APIErrorCode, message string) {
	c.JSON(status, errorResponse{
		Code:  code,
		Error: message,
	})
}
