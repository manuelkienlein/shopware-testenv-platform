package images

import (
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
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

	ctx := c.Request().Context()

	// Get docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// List images
	images, err := cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		panic(err)
	}

	var output []Image

	for _, image := range images {

		// Extract id
		imageHash := strings.Split(image.ID, ":")[1]

		// Extract image name (only uses first)
		imageName := strings.Split(image.RepoTags[0], ":")[0]

		// Extract image tag (only uses first)
		imageTag := strings.Split(image.RepoTags[0], ":")[1]

		output = append(output, Image{
			ID:        imageHash,
			ImageName: imageName,
			ImageTag:  imageTag,
			CreatedAt: time.Unix(image.Created, 0).Format(time.RFC3339),
			Size:      image.Size,
		})
	}

	return c.JSON(http.StatusOK, output)
}
