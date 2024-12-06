package handler

import (
	"context"
	"fmt"
	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

type ContainerInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
	State     string `json:"state"`
	Status    string `json:"status"`
}

// ListContainers listet alle Docker-Container auf
func ListSandboxesHandler(c echo.Context) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Found %d containers\n", len(containers))
	for _, container := range containers {
		fmt.Printf("ID: %s, Image: %s, Status: %s\n", container.ID, container.Image, container.Status)
	}

	var containerInfos []ContainerInfo

	for _, container := range containers {
		// Hole die Startzeit des Containers
		created := time.Unix(container.Created, 0).Format(time.RFC3339)

		// Erstelle ein ContainerInfo Objekt mit den gewünschten Feldern
		containerInfo := ContainerInfo{
			ID:        container.ID,
			Name:      container.Names[0], // Verwende den ersten Namen des Containers
			Image:     container.Image,
			CreatedAt: created,
			State:     container.State,
			Status:    container.Status,
		}

		// Füge das ContainerInfo Objekt zur Liste hinzu
		containerInfos = append(containerInfos, containerInfo)
	}

	// JSON-Antwort zurückgeben
	return c.JSON(http.StatusOK, containerInfos)
}

func CreateSandboxHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

func DeleteSandboxHandler(c echo.Context) error {

	id := c.Param("id")
	log.Println("Stopping sandbox container " + id)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, containertypes.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.Names[0])
		// TODO: Only delete container based on ID
		if container.Names[0] == "/shopware_6.6" {
			fmt.Print("Stopping container ", container.ID[:10], "... ")
			noWaitTimeout := 0 // to not wait for the container to exit gracefully
			if err := cli.ContainerStop(ctx, container.ID, containertypes.StopOptions{Timeout: &noWaitTimeout}); err != nil {
				panic(err)
			}
			fmt.Println("Success")
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{})
}

/*func GetUserHandler(c echo.Context) error {
	// Hole die ID aus den Parametern
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Benutzer finden
	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}*/
