package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/shopwareLabs/testenv-platform/services"
	"log"
	"net/http"
)

type SandboxHandler struct {
	DockerService *services.DockerService
}

func NewSandboxHandler(dockerService *services.DockerService) *SandboxHandler {
	return &SandboxHandler{
		DockerService: dockerService,
	}
}

type ContainerInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	State     string `json:"state"`
	Status    string `json:"status"`
}

// List all sandbox environments
// @Summary List sandbox environments
// @Description List sandbox environments
// @Tags Sandbox Management
// @Accept  json
// @Produce  json
// @Success 200 {object} ContainerInfo
// @Router /api/sandboxes [get]
func (h *SandboxHandler) ListSandboxesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	containerInfos, err := h.DockerService.ListContainers(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, containerInfos)
}

func (h *SandboxHandler) CreateSandboxHandler(c echo.Context) error {
	ctx := c.Request().Context()

	// Image und andere Daten
	imageName := "dockware/dev:6.6.8.2"
	instanceName := "randomString"
	host := "randomString.sandboxes.localhost"

	containerID, err := h.DockerService.CreateContainer(ctx, imageName, instanceName, host)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	log.Printf("Created container with ID: %s", containerID)
	return c.JSON(http.StatusOK, map[string]interface{}{"container_id": containerID})
}

func (h *SandboxHandler) DeleteSandboxHandler(c echo.Context) error {

	ctx := c.Request().Context()

	containerID := c.Param("id")
	log.Println("Stopping sandbox container " + containerID)

	containerInfos, err := h.DockerService.ListContainers(ctx)
	if err != nil {
		panic(err)
	}

	for _, container := range containerInfos {
		fmt.Println(container.Name)
		// TODO: Only delete container based on ID
		if container.Name == "/shopware_6.6" {
			fmt.Print("Stopping container ", container.ID[:10], "... ")

			if err := h.DockerService.StopContainer(ctx, container.ID); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			fmt.Println("Success")
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

/*func GetUserHandler(c echo.Context) error {
	// Hole die ID aus den Parametern
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Benutzer finden
	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}*/
