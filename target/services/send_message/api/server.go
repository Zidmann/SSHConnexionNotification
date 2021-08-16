package api

import (
	"fmt"
	"os"
	"time"

	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

// StartServer starts the server
func StartServer() error {
	// Define the port (8080 by default)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create an Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 10 * time.Second,
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define API group
	g := e.Group("/api")

	/* Unauthenticated route */
	// Health check route
	g.GET("/isalive", isAlive)

	/* Authenticated route */

	// Extract the jwt signing key from the secret vault
	jwt_secret := getJwtKey()

	// Configure middleware with the custom claims type
	config := getJWTconfig(jwt_secret)
	g.Use(middleware.JWTWithConfig(config))

	g.POST("/message/:channel", sendMessage)

	// Server
	return e.Start(fmt.Sprintf(":%s", port))
}
