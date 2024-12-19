package sandboxes

import (
	"github.com/shopwareLabs/testenv-platform/services/sandbox"
)

type SandboxHandler struct {
	SandboxService *sandbox.SandboxService
}

func NewSandboxHandler(sandboxService *sandbox.SandboxService) *SandboxHandler {
	return &SandboxHandler{
		SandboxService: sandboxService,
	}
}
