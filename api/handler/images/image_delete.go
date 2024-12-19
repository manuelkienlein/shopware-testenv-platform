package images

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type ImageDeleteResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Docker Image removed successfully"`
}

// ImageDeleteHandler removes a docker image
// @Summary Remove Docker Image
// @Description Removes a docker image form the system
// @Tags Docker Image Management
// @Accept json
// @Produce json
// @Param id path string true "Image ID"
// @Success 200 {object} ImageDeleteResponse
// @Failure 400 {object} map[string]string
// @Router /api/images/{id} [delete]
func (h *ImageHandler) ImageDeleteHandler(c echo.Context) error {

	containerID := c.Param("id")

	// TODO implement
	output := ImageDeleteResponse{
		Message: "Image " + containerID + " removed successfully",
		Status:  "success",
	}

	return c.JSON(http.StatusOK, output)
}
