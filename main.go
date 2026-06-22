package main

import (
	"context"
	"ctrl-hub-technical-challenge/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := httpserver.NewHttpServer()

	// Run the server in the background so main can wait for a shutdown signal.
	serveErr := make(chan error, 1)
	go func() {
		serveErr <- server.Serve()
	}()

	// Block until an interrupt/termination signal arrives or the server stops
	// on its own with an error.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serveErr:
		if err != nil {
			log.Fatalf("server error: %v", err)
		}
		return
	case <-ctx.Done():
		log.Println("shutdown signal received, stopping server")
	}

	// Allow in-flight requests up to 10s to complete before forcing shutdown.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("server stopped")
}
