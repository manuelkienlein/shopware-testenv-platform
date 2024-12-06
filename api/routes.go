package api

import (
	"github.com/labstack/echo/v4"
	"github.com/shopwareLabs/testenv-platform/api/handler"
)

func RegisterRoutes(e *echo.Echo) {
	// API-Gruppe
	api := e.Group("/api")

	api.GET("/health", handler.HealthCheckHandler)

	api.GET("/sandboxes", handler.ListSandboxesHandler)
	api.POST("/sandboxes", handler.CreateSandboxHandler)

	/*// Beispielroute: GET /api/health
	api.GET("/health", handlers.HealthCheckHandler)

	// Beispielroute: GET /api/users/:id
	api.GET("/users/:id", handlers.GetUserHandler)

	// Beispielroute: POST /api/users
	api.POST("/users", handlers.CreateUserHandler)*/
}
