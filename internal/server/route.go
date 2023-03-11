package server

import "github.com/gin-gonic/gin"

// Route represents a Server route
type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
