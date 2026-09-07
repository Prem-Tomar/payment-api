package main

import (
	"context"
	"fmt"
	"log"	
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Prem-Tomar/payment-api/internal/httpapi"
	"github.com/Prem-Tomar/payment-api/internal/logging"
)


func main() {

	logger := logging.New(os.Stdout)
	router := httpapi.NewRouter(logger)

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
	shutdown, stop := context.WithTimeout(context.Background(), 
	10 * time.Second) 

	defer stop()
	<-shutdown.Done()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
		return
	}

	fmt.Println("Shutdown complete")
	fmt.Println("Shutdown completee")
}

func serverRunner(server *http.Server) {
	if err := server.ListenAndServe(); err != nil &&
	err != http.ErrServerClosed {
	log.Printf("G2 Gin server failed: %v", err)
}
}