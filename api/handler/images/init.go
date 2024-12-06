package images

import "github.com/shopwareLabs/testenv-platform/services"

type ImageHandler struct {
	DockerService *services.DockerService
}

func NewImageHandler(dockerService *services.DockerService) *ImageHandler {
	return &ImageHandler{
		DockerService: dockerService,
	}
}
