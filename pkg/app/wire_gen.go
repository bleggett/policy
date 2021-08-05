// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"github.com/aserto-dev/policy-cli/pkg/cc"
	"github.com/aserto-dev/policy-cli/pkg/cc/config"
	"github.com/google/wire"
	"io"
)

// Injectors from wire.go:

func BuildPolicyCLI(logWriter io.Writer, configPath config.Path, overrides config.Overrider) (*PolicyCLI, func(), error) {
	ccCC, cleanup, err := cc.NewCC(logWriter, configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	context := ccCC.Context
	logger := ccCC.Log
	configConfig := ccCC.Config
	policyCLI := &PolicyCLI{
		Context:       context,
		Logger:        logger,
		Configuration: configConfig,
	}
	return policyCLI, func() {
		cleanup()
	}, nil
}

func BuildTestPolicyCLI(logWriter io.Writer, configPath config.Path, overrides config.Overrider) (*PolicyCLI, func(), error) {
	ccCC, cleanup, err := cc.NewTestCC(logWriter, configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	context := ccCC.Context
	logger := ccCC.Log
	configConfig := ccCC.Config
	policyCLI := &PolicyCLI{
		Context:       context,
		Logger:        logger,
		Configuration: configConfig,
	}
	return policyCLI, func() {
		cleanup()
	}, nil
}

// wire.go:

var (
	policyCLISet = wire.NewSet(cc.NewCC, wire.FieldsOf(new(*cc.CC), "Config", "Log", "Context", "ErrGroup"))

	policyCLITestSet = wire.NewSet(cc.NewTestCC, wire.FieldsOf(new(*cc.CC), "Config", "Log", "Context", "ErrGroup"))
)
