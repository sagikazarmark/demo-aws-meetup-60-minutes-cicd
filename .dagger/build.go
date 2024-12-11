package main

import "github.com/sagikazarmark/demo-aws-meetup-60-minutes-cicd/.dagger/internal/dagger"

func (m *Demo) Build() *dagger.Container {
	binary := dag.Go().
		WithCgoDisabled().
		Build(m.Source, dagger.GoBuildOpts{
            Platform: dagger.Platform("linux/amd64"),
        })

	container := dag.Container().
		From("alpine").
		WithFile("/usr/local/bin/demo", binary).
		WithEntrypoint([]string{"/usr/local/bin/demo"})

	return container
}
