package main

import (
    "context"
    "errors"
    "fmt"
    "time"
)

// Release a new version of the application and push it to the container registry.
func (m *Demo) Release(
    ctx context.Context,

    // +optional
    version string,
) (string, error) {
    if m.Aws == nil {
        return "", errors.New("setup requires AWS to be configured")
    }

    // Authenticate with the registry
    registry := m.Aws.Ecr().Login()

    registryAddress, err := registry.Address(ctx)
    if err != nil {
        return "", err
    }

    if version == "" {
        version = fmt.Sprintf("%d", time.Now().Unix())
    }

    repositoryAddress := registryAddress + "/" + repositoryName + ":" + version

    // Build and push an initial container
    _, err = m.Build().
        With(registry.Auth).
        Publish(ctx, repositoryAddress)
    if err != nil {
        return "", err
    }

	return repositoryAddress, nil
}
