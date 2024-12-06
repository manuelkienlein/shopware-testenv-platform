package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shopwareLabs/testenv-platform/api/handler"
	"github.com/shopwareLabs/testenv-platform/services"
)

func RegisterRoutes(e *echo.Echo) {

	// Init services
	dockerService, err := services.NewDockerService()
	if err != nil {
		e.Logger.Fatalf("Failed to create Docker service: %v", err)
	}

	// Init handlers
	sandboxHandler := handler.NewSandboxHandler(dockerService)

	// Add api handlers
	api := e.Group("/api")

	api.GET("/health", handler.HealthCheckHandler)

	api.GET("/sandboxes", sandboxHandler.ListSandboxesHandler)
	api.POST("/sandboxes", sandboxHandler.CreateSandboxHandler)
	api.DELETE("/sandboxes/:id", sandboxHandler.DeleteSandboxHandler)
}
