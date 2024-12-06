package images

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// ImageListHandler lists all images
// @Summary List of Docker Image
// @Description Get a list of docker image
// @Tags Docker Image Management
// @Accept json
// @Produce json
// @Success 200 {array} Image
// @Failure 400 {object} map[string]string
// @Router /api/images [get]
func (h *ImageHandler) ImageListHandler(c echo.Context) error {

	// TODO implement

	output := []Image{{
		ID:        "containerID",
		ImageName: "Name",
		ImageTag:  "6.0.0.0",
		CreatedAt: "2013-08-20T18:52:09.000Z",
		Size:      123,
	}}

	return c.JSON(http.StatusOK, output)
}
