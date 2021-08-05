// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package cc

import (
	"github.com/aserto-dev/go-lib/certs"
	"github.com/aserto-dev/go-lib/logger"
	"github.com/aserto-dev/policy-cli/pkg/cc/config"
	"github.com/aserto-dev/policy-cli/pkg/cc/context"
	"github.com/google/wire"
	"io"
)

// Injectors from wire.go:

// buildCC sets up the CC struct that contains all dependencies that
// are cross cutting
func buildCC(logOutput io.Writer, configPath config.Path, overrides config.Overrider) (*CC, func(), error) {
	errGroupAndContext := context.NewContext()
	contextContext := errGroupAndContext.Ctx
	loggerConfig, err := config.NewLoggerConfig(configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	zerologLogger, err := logger.NewLogger(logOutput, loggerConfig)
	if err != nil {
		return nil, nil, err
	}
	generator := certs.NewGenerator(zerologLogger)
	configConfig, err := config.NewConfig(configPath, zerologLogger, overrides, generator)
	if err != nil {
		return nil, nil, err
	}
	group := errGroupAndContext.ErrGroup
	ccCC := &CC{
		Context:  contextContext,
		Config:   configConfig,
		Log:      zerologLogger,
		ErrGroup: group,
	}
	return ccCC, func() {
	}, nil
}

func buildTestCC(logOutput io.Writer, configPath config.Path, overrides config.Overrider) (*CC, func(), error) {
	errGroupAndContext := context.NewTestContext()
	contextContext := errGroupAndContext.Ctx
	loggerConfig, err := config.NewLoggerConfig(configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	zerologLogger, err := logger.NewLogger(logOutput, loggerConfig)
	if err != nil {
		return nil, nil, err
	}
	generator := certs.NewGenerator(zerologLogger)
	configConfig, err := config.NewConfig(configPath, zerologLogger, overrides, generator)
	if err != nil {
		return nil, nil, err
	}
	group := errGroupAndContext.ErrGroup
	ccCC := &CC{
		Context:  contextContext,
		Config:   configConfig,
		Log:      zerologLogger,
		ErrGroup: group,
	}
	return ccCC, func() {
	}, nil
}

// wire.go:

var (
	ccSet = wire.NewSet(context.NewContext, config.NewConfig, config.NewLoggerConfig, logger.NewLogger, certs.NewGenerator, wire.FieldsOf(new(config.Config), "Logging"), wire.FieldsOf(new(*context.ErrGroupAndContext), "Ctx", "ErrGroup"), wire.Struct(new(CC), "*"))

	ccTestSet = wire.NewSet(context.NewTestContext, config.NewConfig, config.NewLoggerConfig, logger.NewLogger, certs.NewGenerator, wire.FieldsOf(new(*context.ErrGroupAndContext), "Ctx", "ErrGroup"), wire.Struct(new(CC), "*"))
)
