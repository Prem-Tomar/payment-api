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

	// Channel used to report server errors back to main
	serverErr := make(chan error , 1)

	go func(){
		serverErr <-serverRunner(server)
	}()


	// creating signals
	shutdown, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	select {
	case err := <-serverErr:
		log.Printf("Server failed: %v", err)
		return

	case <-shutdown.Done():
		fmt.Println("Shutting down server")
	}

	// Give active requests a maximum of 10 seconds to finish.
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown failed: %v", err)
		return
	}

	fmt.Println("Shutdown completee")
}

func serverRunner(server *http.Server) error {
	if err := server.ListenAndServe(); err != nil &&
	err != http.ErrServerClosed {
	log.Printf("G2 Gin server failed: %v", err)
	return err
}
return nil
}