package main

import (
    "context"
    "errors"
)

// Set up the initial environment.
func (m *Demo) Setup(ctx context.Context) error {
    if m.Aws == nil {
        return errors.New("setup requires AWS to be configured")
    }

    // Create the ECR repository
    _, err := m.Aws.Ecr().Exec([]string{"create-repository", "--repository-name", repositoryName}).Sync(ctx)
    if err != nil {
        return err
    }

    // Release an initial version
    repositoryAddress, err := m.Release(ctx, "latest")

    sourceConfiguration, err := m.sourceConfiguration(ctx, repositoryAddress)
    if err != nil {
        return err
    }

    // Create the initial AppRunner service
    _, err = m.Aws.Exec([]string{
        "apprunner", "create-service",
        "--service-name", serviceName,
        "--source-configuration", sourceConfiguration,
    }).Sync(ctx)
    if err != nil {
        return err
    }

	return nil
}

// Tear down all resources.
func (m *Demo) Teardown(ctx context.Context) error {
    if m.Aws == nil {
        return errors.New("setup requires AWS to be configured")
    }

    // Delete the ECR repository
    _, err := m.Aws.Ecr().Exec([]string{"delete-repository", "--force", "--repository-name", repositoryName}).Sync(ctx)
    if err != nil {
        return err
    }

    serviceArn, err := m.getServiceArn(ctx)
    if err != nil {
        return err
    }

    // Delete the AppRunner service
    _, err = m.Aws.Exec([]string{
        "apprunner", "delete-service",
        "--service-arn", serviceArn,
    }).Sync(ctx)
    if err != nil {
        return err
    }

	return nil
}
