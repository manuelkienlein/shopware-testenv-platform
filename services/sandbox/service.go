package sandbox

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"time"
)

type SandboxService struct {
	database *SandboxDatabase
	client   *client.Client
}

func NewSandboxService() (*SandboxService, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &SandboxService{
		database: NewSandboxDatabase(),
		client:   cli,
	}, nil
}

func (s *SandboxService) ListSandboxes(ctx context.Context) ([]SandboxInfo, error) {
	containers, err := s.client.ContainerList(ctx, container.ListOptions{})
	if err != nil {
		return nil, err
	}

	sandboxInfos := make([]SandboxInfo, 0)
	for _, cont := range containers {
		created := time.Unix(cont.Created, 0).Format(time.RFC3339)

		if cont.Labels["sandbox_container"] != "true" {
			continue
		}

		containerInfo := SandboxInfo{
			ID:            cont.Labels["sandbox_id"],
			ContainerName: cont.Names[0],
			ContainerId:   cont.ID,
			Url:           cont.Labels["sandbox_host"],
			Image:         cont.Image,
			CreatedAt:     created,
			State:         cont.State,
			Status:        cont.Status,
		}
		sandboxInfos = append(sandboxInfos, containerInfo)
	}

	return sandboxInfos, nil
}

func (s *SandboxService) CreateSandbox(ctx context.Context, imageName string) (Sandbox, error) {

	sandboxId := uuid.New().String()
	containerName := "sandbox-" + sandboxId
	hostname := containerName + ".shopshredder.zion.mr-pixel.de"

	// Pull docker container
	out, err := s.client.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		log.Fatal("Failed to pull sandbox docker container", err)
		return Sandbox{}, err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	// Create docker container
	labels := map[string]string{
		"sandbox_container": "true",
		"sandbox_id":        sandboxId,
		"sandbox_host":      hostname,
		"traefik.enable":    "true",
		fmt.Sprintf("traefik.http.routers.http-%s.rule", containerName): fmt.Sprintf("Host(`%s`)", hostname),
	}
	resp, err := s.client.ContainerCreate(ctx, &container.Config{
		Image:  imageName,
		Labels: labels,
	}, nil, nil, nil, containerName)
	if err != nil {
		log.Fatal("Failed to create sandbox docker container", err)
		return Sandbox{}, err
	}

	// Start docker container
	if err := s.client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatal("Failed to start sandbox docker container", err)
		return Sandbox{}, err
	}
	log.Printf("Started sandbox %s with image %s", containerName, imageName)

	// Register sandbox in database
	sandbox := Sandbox{
		ID:            sandboxId,
		ImageName:     imageName,
		ImageId:       "",
		ContainerName: containerName,
		ContainerId:   resp.ID,
		Url:           "https://" + hostname,
		CreatedAt:     time.Now(),
		LifeTime:      1440,
	}
	s.database.Add(sandbox)

	return sandbox, nil
}

func (s *SandboxService) DeleteSandbox(ctx context.Context, sandboxId string) {

	// Find containerId for sandboxId
	sandbox, err := s.database.GetByID(sandboxId)
	if err != nil {
		log.Fatalf("Failed to delete sandbox %s, because sandbox not found: %v", sandboxId, err)
	}

	// Stop sandbox container
	noWaitTimeout := 0 // to not wait for the container to exit gracefully
	err = s.client.ContainerStop(ctx, sandbox.ContainerId, container.StopOptions{Timeout: &noWaitTimeout})
	if err != nil {
		log.Fatalf("Failed to stop sandbox container %s: %v", sandbox.ContainerName, err)
	}

	// Remove sandbox from database
	err = s.database.Remove(sandboxId)
	if err != nil {
		log.Fatalf("Failed to delete sandbox %s in database: %v", sandboxId, err)
	}
}

func (s *SandboxService) ShutdownSandboxes() {
	// TODO
}

type SandboxInfo struct {
	ID            string `json:"id"`
	ContainerId   string `json:"container_id"`
	ContainerName string `json:"container_name"`
	Url           string `json:"url"`
	Image         string `json:"image"`
	CreatedAt     string `json:"created_at"`
	State         string `json:"state"`
	Status        string `json:"status"`
}
