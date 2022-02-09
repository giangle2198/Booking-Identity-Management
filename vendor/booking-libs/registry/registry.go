package registry

import (
	"sync"

	"github.com/sarulabs/di"
)

type DIBuilder func() []di.Def

var (
	buildOnce sync.Once
	builder   *di.Builder
	container di.Container
	// ConfigdBuilder builder for config
	ConfigsBuilder DIBuilder
	// HelpersBuilder builder for all helpers
	HelpersBuilder DIBuilder
	// RepositoriesBuilder builder for repositories
	RepositoriesBuilder DIBuilder
	// AdaptersBuilder builder for adapters
	AdaptersBuilder DIBuilder
	// UsecasesBuilder builder for usecase
	UsecasesBuilder DIBuilder
	// APIsBuilder builder for apis
	APIsBuilder DIBuilder
)

func BuildDIContainer() {
	buildOnce.Do(func() {
		builder, _ = di.NewBuilder()
		doBuild()
		container = builder.Build()
	})
}

func doBuild() {
	if err := buildConfigs(); err != nil {
		panic(err)
	}
	if err := buildHelpers(); err != nil {
		panic(err)
	}
	if err := buildRepositories(); err != nil {
		panic(err)
	}
	if err := buildAdapters(); err != nil {
		panic(err)
	}
	if err := buildUsecases(); err != nil {
		panic(err)
	}
	if err := buildAPIs(); err != nil {
		panic(err)
	}
}

func GetDependency(dependencyName string) interface{} {
	return container.Get(dependencyName)
}

func CleanDependency() error {
	return container.Clean()
}

func buildConfigs() error {
	if ConfigsBuilder == nil {
		ConfigsBuilder = defaultBuilder
	}
	defs := ConfigsBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildAPIs() error {
	if APIsBuilder == nil {
		APIsBuilder = defaultBuilder
	}
	defs := APIsBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildHelpers() error {
	if HelpersBuilder == nil {
		HelpersBuilder = defaultBuilder
	}
	defs := HelpersBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildUsecases() error {
	if UsecasesBuilder == nil {
		UsecasesBuilder = defaultBuilder
	}
	defs := UsecasesBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildRepositories() error {
	if RepositoriesBuilder == nil {
		RepositoriesBuilder = defaultBuilder
	}
	defs := RepositoriesBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func buildAdapters() error {
	if AdaptersBuilder == nil {
		AdaptersBuilder = defaultBuilder
	}
	defs := AdaptersBuilder()
	err := builder.Add(defs...)
	if err != nil {
		return err
	}
	return nil
}

func defaultBuilder() []di.Def {
	return []di.Def{}
}
