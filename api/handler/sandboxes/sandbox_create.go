package sandboxes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SandboxCreateRequest struct {
	ImageId  string `json:"image_id" example:"dockware/dev"`
	Lifetime int    `json:"lifetime" example:"1440"`
}

type SandboxCreateResponse struct {
	Status        string `json:"status" example:"success"`
	Message       string `json:"message" example:"Sandbox created successfully"`
	Image         string `json:"image" example:"dockware/dev:6.6.8.2"`
	ContainerName string `json:"container_name" example:"sandbox-container"`
	ContainerId   string `json:"container_id" example:"sandbox-container-id"`
	Url           string `json:"url" example:"https://sandbox-random.shopshredder.zion.mr-pixel.de"`
}

// SandboxCreateHandler creates a new sandbox for requested image
// @Summary Create new sandbox
// @Description Creates a new sandbox docker container
// @Tags Sandbox Management
// @Accept json
// @Produce json
// @Param image body SandboxCreateRequest true "Image Input"
// @Success 200 {object} SandboxCreateResponse
// @Failure 400 {object} map[string]string
// @Router /api/sandboxes [post]
func (h *SandboxHandler) SandboxCreateHandler(c echo.Context) error {

	// TODO implement
	ctx := c.Request().Context()
	var input SandboxCreateRequest

	// Parse input
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request format",
		})
	}

	// Input validation
	if input.Lifetime > 1440 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Lifetime should be less than 1440",
		})
	}

	if input.Lifetime < 5 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Lifetime should be greater than 5",
		})
	}

	// TODO: other data
	imageName := "dockware/dev:6.6.8.2"
	instanceName := "sandbox_random"
	host := "sandbox-random.shopshredder.zion.mr-pixel.de"

	containerID, err := h.DockerService.CreateContainer(ctx, imageName, instanceName, host)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	output := SandboxCreateResponse{
		Message:       "Sandbox created successfully",
		Status:        "success",
		ContainerName: instanceName,
		ContainerId:   containerID,
		Image:         imageName,
		Url:           host,
	}

	return c.JSON(http.StatusOK, output)
}
