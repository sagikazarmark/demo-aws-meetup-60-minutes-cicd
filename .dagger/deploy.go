package main

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

const sourceConfigurationTpl = `
{
    "ImageRepository": {
        "ImageIdentifier": "%s",
        "ImageRepositoryType": "ECR"
    },
    "AutoDeploymentsEnabled": true,
    "AuthenticationConfiguration": {
        "AccessRoleArn": "arn:aws:iam::%s:role/AppRunnerECRAccessRole"
    }
}
`

func (m *Demo) sourceConfiguration(ctx context.Context, repository string) (string, error) {
	// Get the AWS account ID
	accountId, err := m.Aws.Sts().GetCallerIdentity().Account(ctx)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(sourceConfigurationTpl, repository, accountId), nil
}

// Get the AppRunner service ARN
func (m *Demo) getServiceArn(ctx context.Context) (string, error) {
	serviceArn, err := m.Aws.Exec([]string{
		"apprunner", "list-services",
		"--query", fmt.Sprintf(`ServiceSummaryList[?ServiceName == '%s'].ServiceArn`, serviceName),
		"--output", "text",
	}).Stdout(ctx)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(serviceArn), nil
}

// Deploy a new version of the application.
func (m *Demo) Deploy(
	ctx context.Context,

	// +optional
	version string,
) error {
	if m.Aws == nil {
		return errors.New("setup requires AWS to be configured")
	}

	// Release a new version
	repositoryAddress, err := m.Release(ctx, version)

	serviceArn, err := m.getServiceArn(ctx)
	if err != nil {
		return err
	}

	sourceConfiguration, err := m.sourceConfiguration(ctx, repositoryAddress)
	if err != nil {
		return err
	}

	// Update the AppRunner service
	_, err = m.Aws.Exec([]string{
		"apprunner", "update-service",
		"--service-arn", serviceArn,
		"--source-configuration", sourceConfiguration,
	}).Sync(ctx)
	if err != nil {
		return err
	}

	return nil
}
