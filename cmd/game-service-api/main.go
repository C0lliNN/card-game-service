package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := newServer()
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			log.Default().Fatalf("failed starting server: %v", err)
		}
	}()
	log.Println("starting HTTP server")

	<-done
	log.Println("shutting down HTTP server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Default().Fatalf("failed to shutdown HTTP server")
	}

	log.Println("server was shutdown properly")
}
