package sandboxes

import "github.com/shopwareLabs/testenv-platform/services"

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
