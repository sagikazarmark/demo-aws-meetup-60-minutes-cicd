package main

import (
	"github.com/sagikazarmark/demo-aws-meetup-60-minutes-cicd/.dagger/internal/dagger"
)

type Demo struct {
	// Project source directory
	//
	// +private
	Source *dagger.Directory

	// +private
    Aws *dagger.AwsCli
}

func New(
	// Project source directory.
	//
	// +defaultPath="/"
	// +ignore=[".devenv", ".direnv", ".github", ".vscode", "go.work", "go.work.sum"]
	source *dagger.Directory,

    // +default="eu-central-1"
    region string,

    // +optional
    awsAccessKeyId *dagger.Secret,

    // +optional
    awsSecretAccessKey *dagger.Secret,
) *Demo {
    var aws *dagger.AwsCli

    if awsAccessKeyId != nil && awsSecretAccessKey != nil {
        aws = dag.AwsCli().
            WithRegion(region).
            WithStaticCredentials(awsAccessKeyId, awsSecretAccessKey)
    }

	return &Demo{
		Source: source,
        Aws: aws,
	}
}
