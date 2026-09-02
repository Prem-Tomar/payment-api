package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {

	router := gin.Default()

	router.GET("/get" , getRequesthandler)

	if err := router.Run() ; err != nil {
		log.Fatal("G2 Gin server failed : %v " , err)
	}
}

func getRequesthandler(context *gin.Context){

	context.JSON(http.StatusOK , gin.H{
		"Statuf_Gin_Server" : "Server Started",
	})
}