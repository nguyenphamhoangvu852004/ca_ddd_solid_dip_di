package initialize

import (
	"net/http"

	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApplication() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Use default CORS config (allow anywhere to communicate with server)
	r.Use(cors.Default())

	// Use default Helmet config (set some security into HTTP response headers)
	r.Use(ginhelmet.Default())

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	routerGroup := r.Group("/api/v1")

	registerRouters(routerGroup, InitDependencies())

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
