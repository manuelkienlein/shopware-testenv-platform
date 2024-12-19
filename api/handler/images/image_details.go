package images

import (
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Image struct {
	ID        string `json:"id" example:"a407dee395ed97ead1e40c7537395d6271c07cc89c317f8eda1c19f6fc783695"`
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
// @Param id path string true "Image ID" example(a407dee395ed97ead1e40c7537395d6271c07cc89c317f8eda1c19f6fc783695)
// @Success 200 {object} Image
// @Failure 400 {object} map[string]string
// @Router /api/images/{id} [get]
func (h *ImageHandler) ImageDetailsHandler(c echo.Context) error {

	ctx := c.Request().Context()
	imageId := c.Param("id")

	// Get docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Get image
	image, _, err := cli.ImageInspectWithRaw(ctx, imageId)
	if err != nil {
		panic(err)
	}

	var output []Image

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
		CreatedAt: image.Created,
		Size:      image.Size,
	})

	return c.JSON(http.StatusOK, output)
}
