package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/c0x12c/backend-go/internal/route"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create Echo instance
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	// Add middleware
	server.Use(echoMiddleware.Logger())
	server.Use(echoMiddleware.Recover())
	server.Use(echoMiddleware.CORS())

	// Configure routes
	healthRouter := route.NewHealthRouter()
	healthRouter.Configure(server)

	// Start server
	port := getPort()
	go func() {
		if err := server.Start(fmt.Sprintf(":%s", port)); err != nil {
			server.Logger.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	server.Logger.Info("Shutting down server...")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
