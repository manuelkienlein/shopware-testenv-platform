package images

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Image struct {
	ID        string `json:"id" example:"a407dee395ed"`
	ImageName string `json:"image_name" example:"dockware/dev"`
	ImageTag  string `json:"image_tag" example:"6.6.8.2"`
	CreatedAt string `json:"created_at" example:"2013-08-20T18:52:09.000Z"`
	Size      int64  `json:"size" example:"1048576"`
}

// ImageDetailsHandler returns information about a image
// @Summary Details about a Docker Image
// @Description Get details about a docker image
// @Tags Docker Image Management
// @Accept json
// @Produce json
// @Param id path string true "Image ID"
// @Success 200 {object} Image
// @Failure 400 {object} map[string]string
// @Router /api/images/{id} [get]
func (h *ImageHandler) ImageDetailsHandler(c echo.Context) error {

	containerID := c.Param("id")

	// TODO implement

	output := Image{
		ID:        containerID,
		ImageName: "Name",
		ImageTag:  "6.0.0.0",
		CreatedAt: "2013-08-20T18:52:09.000Z",
		Size:      123,
	}

	return c.JSON(http.StatusOK, output)
}
