package main

import (
	
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/Prem-Tomar/payment-api/internal/httpapi"
)


func main() {
	
	router := httpapi.NewRouter()

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
}



func serverRunner(server *http.Server) {
	if err := server.ListenAndServe(); err != nil &&
	err != http.ErrServerClosed {
	log.Printf("G2 Gin server failed: %v", err)
}
}