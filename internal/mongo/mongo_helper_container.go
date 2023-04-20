package mongo

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var mongoMux = new(sync.Mutex)

const (
	defaultMongoDBPort = 27017
)

type TestContainer struct {
	testcontainers.Container
	URI string
}

// StartMongoContainer starts a mongodb container using testcontainers-go
func StartMongoContainer(ctx context.Context) (*TestContainer, error) {
	mongoMux.Lock()
	defer mongoMux.Unlock()
	mongoPort, _ := nat.NewPort("", strconv.Itoa(defaultMongoDBPort))

	timeout := 5 * time.Minute

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{string(mongoPort) + "/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(timeout),
		AutoRemove:   true,
	}

	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		})
	if err != nil {
		err = fmt.Errorf("failed to start mongo container: %w", err)
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get mongo container ip: %w", err)
		return nil, err
	}

	port, err := container.MappedPort(ctx, mongoPort)
	if err != nil {
		err = fmt.Errorf("failed to get exposed mongo container port: %w", err)
		return nil, err
	}

	return &TestContainer{
		Container: container,
		URI:       getContainerURI(ip, port),
	}, nil
}

func getContainerURI(host string, port nat.Port) string {
	databaseHost := fmt.Sprintf("%s:%d", host, port.Int())
	return fmt.Sprintf("mongodb://%s", databaseHost)
}
