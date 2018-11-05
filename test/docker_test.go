package test

import (
  "fmt"
  "strconv"
  "testing"
  "time"

  "github.com/gruntwork-io/terratest/modules/docker"
  "github.com/gruntwork-io/terratest/modules/http-helper"
)

func TestPackerDockerExampleLocal(t *testing.T) {
	serverPort := 8080

	dockerOptions := &docker.Options{
		// Directory where docker-compose.yml lives
		WorkingDir: "../docker/python",

		// Configure the port the web app will listen on and the text it will return using environment variables
		EnvVars: map[string]string{
			"SERVER_PORT": strconv.Itoa(serverPort),
		},
	}

	// Make sure to shut down the Docker container at the end of the test
	defer docker.RunDockerCompose(t, dockerOptions, "down")

	// Run Docker Compose to fire up the web app. We run it in the background (-d) so it doesn't block this test.
	docker.RunDockerCompose(t, dockerOptions, "up", "-d")

	// It can take a few seconds for the Docker container boot up, so retry a few times
	maxRetries := 5
	timeBetweenRetries := 2 * time.Second
	url := fmt.Sprintf("http://localhost:%d", serverPort)

	// Verify that we get back a 200 OK with the expected text
	http_helper.HttpGetWithRetry(t, url, 200, "Hello World", maxRetries, timeBetweenRetries)
}
