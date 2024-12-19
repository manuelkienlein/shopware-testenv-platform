package sandboxes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type SandboxDeleteResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Docker Image removed successfully"`
}

// SandboxDeleteHandler removes a docker sandbox
// @Summary Remove Sandbox Environment
// @Description Removes a sandbox docker container
// @Tags Sandbox Management
// @Accept json
// @Produce json
// @Param id path string true "Container ID"
// @Success 200 {object} SandboxDeleteResponse
// @Failure 400 {object} map[string]string
// @Router /api/sandboxes/{id} [delete]
func (h *SandboxHandler) SandboxDeleteHandler(c echo.Context) error {

	ctx := c.Request().Context()
	containerID := c.Param("id")

	// TODO implement

	log.Println("Stopping sandbox container " + containerID)

	containerInfos, err := h.DockerService.ListContainers(ctx)
	if err != nil {
		panic(err)
	}

	for _, container := range containerInfos {
		fmt.Println("ContainerName " + container.Name)
		fmt.Println("ContainerID " + container.ID)
		// TODO: Only delete container based on ID
		if container.Name == "/shopware_6.6" {
			fmt.Print("Stopping container ", container.ID[:10], "... ")

			if err := h.DockerService.StopContainer(ctx, container.ID); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}

			fmt.Println("Success")
		}
	}

	output := SandboxDeleteResponse{
		Message: "Sandbox " + containerID + " removed successfully",
		Status:  "success",
	}

	return c.JSON(http.StatusOK, output)
}
