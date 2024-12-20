package images

import (
	"github.com/shopwareLabs/testenv-platform/services"
	"github.com/shopwareLabs/testenv-platform/services/images"
)

type ImageHandler struct {
	DockerService *services.DockerService
	ImageService  *images.ImageService
}

func NewImageHandler(dockerService *services.DockerService, imageService *images.ImageService) *ImageHandler {
	return &ImageHandler{
		DockerService: dockerService,
		ImageService:  imageService,
	}
}
