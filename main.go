package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shopwareLabs/testenv-platform/api"
	"log"
	"net/http"
)

func main() {
	//go handler.PullImageUpdatesTask()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Loggt Anfragen
	e.Use(middleware.Recover()) // Fängt Panics ab und gibt 500 zurück
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://www.shopshredder.de", "http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Access-Control-Allow-Origin"},
	}))

	// Register routes
	api.RegisterRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	port := ":8080"
	log.Printf("Starting server on http://localhost%s", port)
	if err := e.Start(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Could not start server: %v", err)
	}

	/*router := httprouter.New()

	router.GET("/", handler.Info)

	// New Routes
	router.GET("/environments", handler.ListContainer)
	router.POST("/environments", handler.CreateEnvironment)
	router.DELETE("/environments", handler.DeleteContainer)

	log.Println("Go!")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))*/
}
