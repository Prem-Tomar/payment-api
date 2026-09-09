package httpapi

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/text/message"
)

type successResponse struct{
	Status string `json : "status"`
}

type errorResponse struct {
	Error string `type : "error"`
}

func writeSuccess(c *gin.Context, status int , message string){

	c.JSON(status , successResponse{

		Status:  message,
	})

}


func writeError( c *gin.Context , status int , message string){
	c.JSON(status , successResponse{
		Status : message,
	})
}