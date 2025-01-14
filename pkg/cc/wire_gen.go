// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cc

import (
	"github.com/aserto-dev/go-utils/certs"
	"github.com/aserto-dev/go-utils/logger"
	"github.com/google/wire"
	"github.com/bleggett/policy/pkg/cc/config"
	"github.com/bleggett/policy/pkg/cc/context"
)

// Injectors from wire.go:

// buildCC sets up the CC struct that contains all dependencies that
// are cross cutting
func buildCC(logOutput logger.Writer, errOutput logger.ErrWriter, configPath config.Path, overrides config.Overrider) (*CC, func(), error) {
	errGroupAndContext := context.NewContext()
	contextContext := errGroupAndContext.Ctx
	loggerConfig, err := config.NewLoggerConfig(configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	zerologLogger, err := logger.NewLogger(logOutput, errOutput, loggerConfig)
	if err != nil {
		return nil, nil, err
	}
	generator := certs.NewGenerator(zerologLogger)
	configConfig, err := config.NewConfig(configPath, zerologLogger, overrides, generator)
	if err != nil {
		return nil, nil, err
	}
	group := errGroupAndContext.ErrGroup
	cancelFunc := errGroupAndContext.Cancel
	ccCC := &CC{
		Context:    contextContext,
		Config:     configConfig,
		Log:        zerologLogger,
		ErrGroup:   group,
		CancelFunc: cancelFunc,
	}
	return ccCC, func() {
	}, nil
}

func buildTestCC(logOutput logger.Writer, errOutput logger.ErrWriter, configPath config.Path, overrides config.Overrider) (*CC, func(), error) {
	errGroupAndContext := context.NewTestContext()
	contextContext := errGroupAndContext.Ctx
	loggerConfig, err := config.NewLoggerConfig(configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	zerologLogger, err := logger.NewLogger(logOutput, errOutput, loggerConfig)
	if err != nil {
		return nil, nil, err
	}
	generator := certs.NewGenerator(zerologLogger)
	configConfig, err := config.NewConfig(configPath, zerologLogger, overrides, generator)
	if err != nil {
		return nil, nil, err
	}
	group := errGroupAndContext.ErrGroup
	cancelFunc := errGroupAndContext.Cancel
	ccCC := &CC{
		Context:    contextContext,
		Config:     configConfig,
		Log:        zerologLogger,
		ErrGroup:   group,
		CancelFunc: cancelFunc,
	}
	return ccCC, func() {
	}, nil
}

// wire.go:

var (
	ccSet = wire.NewSet(context.NewContext, config.NewConfig, config.NewLoggerConfig, logger.NewLogger, certs.NewGenerator, wire.FieldsOf(new(config.Config), "Logging"), wire.FieldsOf(new(*context.ErrGroupAndContext), "Ctx", "ErrGroup", "Cancel"), wire.Struct(new(CC), "*"))

	ccTestSet = wire.NewSet(context.NewTestContext, config.NewConfig, config.NewLoggerConfig, logger.NewLogger, certs.NewGenerator, wire.FieldsOf(new(*context.ErrGroupAndContext), "Ctx", "ErrGroup", "Cancel"), wire.Struct(new(CC), "*"))
)
