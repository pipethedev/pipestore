package services

import (
	"fmt"
	"os/exec"
)

const (
	imageName     = "ileripipe/pipebase:latest"
	containerName = "pipebase_server"
)

func InstallImageAndRunContainer() error {
	if err := installPipeStoreImage(); err != nil {
		return err
	}

	if isContainerRunning(containerName) {
		fmt.Println("PipeStore container is already running.")
		return nil
	}

	if err := startPipeStoreContainer(); err != nil {
		return err
	}

	return nil
}

func installPipeStoreImage() error {
	cmd := exec.Command("docker", "pull", imageName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error installing PipeStore Docker image:", err)
		fmt.Println(string(output))
		return err
	}

	fmt.Println("PipeStore Docker image installed successfully.")
	return nil
}

func startPipeStoreContainer() error {
	fmt.Println("Starting PipeStore container...")

	cmd := exec.Command("docker", "run", "--name", containerName, "-d", imageName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error starting PipeStore container:", err)
		fmt.Println(string(output))
		return err
	}

	fmt.Println("PipeStore container started successfully.")
	return nil
}

func isContainerRunning(containerName string) bool {
	cmd := exec.Command("docker", "inspect", "-f", "{{.State.Running}}", containerName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error checking container status:", err)
		fmt.Println(string(output))
		return false
	}

	return string(output) == "true"
}
