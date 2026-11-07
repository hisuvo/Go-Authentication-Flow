package main

import (
	"auth-two/internal/config"
	"auth-two/internal/routes"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// Get port from environment.
	port := config.LoadEnv().PORT
	if port == "" {
		port = "8000"
	}

	// Initialize router.
	router := mux.NewRouter()

	// Register application routes.
	routes.RegistersRoutes(router)

	// Configure HTTP server.
	srv := &http.Server{
		Addr:         "127.0.0.1:" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a separate goroutine.
	go func() {
		log.Printf("Server running at http://%s", srv.Addr)

		if err := srv.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Listen for termination signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	// Gracefully shut down the server.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
	
	log.Println("Server stopped successfully")
}
