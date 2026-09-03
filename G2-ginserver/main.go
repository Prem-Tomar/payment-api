package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/get", getRequesthandler)
	router.POST("post", postRequestHandler)
	// to create a explicit router replace router.run with belo code after importing time

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go serverRunner(server) // this is a go routine ans it will carry its work in seperate thread

	// creating signals

	shutdown, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-shutdown.Done()

	fmt.Println("Shutting down server")

	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println("Shutdown with Error -----")
		log.Fatal("Server shutdown failed:", err)

	}

	fmt.Println("Shutdown completee")

	// if err := router.Run() ; err != nil {
	// 	log.Fatal("G2 Gin server failed : %v " , err)
	// }
}

func getRequesthandler(context *gin.Context) {
	time.Sleep(15 * time.Second) // for testing how this timeout thing works
	fmt.Println("server started")
	context.JSON(http.StatusOK, gin.H{
		"Statuf_Gin_Server": "Server Started",
	})
}

func postRequestHandler(context *gin.Context) {

	fmt.Println("Post request is fired")
	type CustomRequest struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var customRequestObject CustomRequest

	context.BindJSON(&customRequestObject)
	context.JSON(http.StatusOK, gin.H{
		"email": customRequestObject.Email,
		"name":  customRequestObject.Name,
	})
}

func serverRunner(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server is shutting down as it is closed noe \n ")
		log.Fatal("G2 Gin server failed : %v ", err)
	}
}
