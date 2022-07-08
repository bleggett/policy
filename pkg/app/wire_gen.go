// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/aserto-dev/clui"
	"github.com/aserto-dev/go-utils/logger"
	"github.com/google/wire"
	"github.com/bleggett/policy/pkg/cc"
	"github.com/bleggett/policy/pkg/cc/config"
)

// Injectors from wire.go:

func BuildPolicyApp(logOutput logger.Writer, errOutput logger.ErrWriter, configPath config.Path, overrides config.Overrider) (*PolicyApp, func(), error) {
	ccCC, cleanup, err := cc.NewCC(logOutput, errOutput, configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	context := ccCC.Context
	cancelFunc := ccCC.CancelFunc
	zerologLogger := ccCC.Log
	configConfig := ccCC.Config
	ui := clui.NewUI()
	policyApp := &PolicyApp{
		Context:       context,
		Cancel:        cancelFunc,
		Logger:        zerologLogger,
		Configuration: configConfig,
		UI:            ui,
	}
	return policyApp, func() {
		cleanup()
	}, nil
}

func BuildTestPolicyApp(logOutput logger.Writer, errOutput logger.ErrWriter, configPath config.Path, overrides config.Overrider) (*PolicyApp, func(), error) {
	ccCC, cleanup, err := cc.NewTestCC(logOutput, errOutput, configPath, overrides)
	if err != nil {
		return nil, nil, err
	}
	context := ccCC.Context
	cancelFunc := ccCC.CancelFunc
	zerologLogger := ccCC.Log
	configConfig := ccCC.Config
	ui := clui.NewUI()
	policyApp := &PolicyApp{
		Context:       context,
		Cancel:        cancelFunc,
		Logger:        zerologLogger,
		Configuration: configConfig,
		UI:            ui,
	}
	return policyApp, func() {
		cleanup()
	}, nil
}

// wire.go:

var (
	policyAppSet = wire.NewSet(cc.NewCC, clui.NewUI, wire.FieldsOf(new(*cc.CC), "Config", "Log", "Context", "ErrGroup", "CancelFunc"))

	policyAppTestSet = wire.NewSet(cc.NewTestCC, clui.NewUI, wire.FieldsOf(new(*cc.CC), "Config", "Log", "Context", "ErrGroup", "CancelFunc"))
)
