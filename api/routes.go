package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shopwareLabs/testenv-platform/api/handler"
	"github.com/shopwareLabs/testenv-platform/api/handler/images"
	_ "github.com/shopwareLabs/testenv-platform/docs"
	"github.com/shopwareLabs/testenv-platform/services"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title mpXsandbox API
// @version 1.0.0
// @description Management API for Docker Sandbox Enviroment
// @license.name MIT
// @host localhost:8080
// @BasePath /api
// @schemes http https
func RegisterRoutes(e *echo.Echo) {

	// Init services
	dockerService, err := services.NewDockerService()
	if err != nil {
		e.Logger.Fatalf("Failed to create Docker service: %v", err)
	}

	// Init handlers
	sandboxHandler := handler.NewSandboxHandler(dockerService)
	imageHandler := images.NewImageHandler(dockerService)

	// Add api handlers
	api := e.Group("/api")

	api.GET("/health", handler.HealthCheckHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api.GET("/sandboxes", sandboxHandler.ListSandboxesHandler)
	api.POST("/sandboxes", sandboxHandler.CreateSandboxHandler)
	api.DELETE("/sandboxes/:id", sandboxHandler.DeleteSandboxHandler)

	api.GET("/images", imageHandler.ImageListHandler)
	api.GET("/images/:id", imageHandler.ImageDetailsHandler)
	api.POST("/images", imageHandler.PullImageHandler)
	api.DELETE("/images/:id", imageHandler.ImageDeleteHandler)
}
