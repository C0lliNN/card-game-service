// Package server implements the code necessary to start and shutdown an HTTP Server

package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Config aggregates configurable parameters for the Server
type Config struct {
	Router          *gin.Engine
	Addr            string
	Timeout         time.Duration
	GameHandler     GameHandler
	ErrorMiddleware ErrorMiddleware
}

// Server wrapper around http.Server
type Server struct {
	Config
	httpServer *http.Server
}

func New(c Config) *Server {
	return &Server{Config: c}
}

// Start sets up and starts the HTTP server
func (s *Server) Start() error {
	router := s.Router

	router.Use(gin.Recovery())
	router.Use(s.ErrorMiddleware.Handler())

	routes := s.GameHandler.Routes()
	for _, r := range routes {
		router.Handle(r.Method, r.Path, r.Handler)
	}

	s.httpServer = &http.Server{
		Handler:      router,
		Addr:         s.Addr,
		WriteTimeout: s.Timeout,
		ReadTimeout:  s.Timeout,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown gracefully shuts down the HTTP server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
