package api

import (
	"errors"
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
	project_id := os.Getenv("PROJECT_ID")
	if project_id == "" {
		return errors.New("No PROJECT_ID environment variable")
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
	g.GET("/alive", alive)

	/* Authenticated route */
	signingKey, err := getJwtSigningKey(project_id)
	if err != nil {
		return err
	}
	config := getJWTconfig(signingKey)
	g.Use(middleware.JWTWithConfig(config))

	// Forward a message
	g.POST("/message", sendMessageFunc(project_id))

	// Server
	return e.Start(fmt.Sprintf(":%s", port))
}
