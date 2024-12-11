package main

import "github.com/sagikazarmark/demo-aws-meetup-60-minutes-cicd/.dagger/internal/dagger"

func (m *Demo) Run() *dagger.Service {
	container := m.Build()

	return container.
		WithExposedPort(8080).
		AsService()
}
