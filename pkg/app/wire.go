//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"

	"github.com/aserto-dev/clui"
	"github.com/aserto-dev/go-utils/logger"
	"github.com/bleggett/policy/pkg/cc"
	"github.com/bleggett/policy/pkg/cc/config"
)

var (
	policyAppSet = wire.NewSet(
		cc.NewCC,
		clui.NewUI,

		wire.FieldsOf(new(*cc.CC), "Config", "Log", "Context", "ErrGroup", "CancelFunc"),
	)

	policyAppTestSet = wire.NewSet(
		// Test
		cc.NewTestCC,

		// Normal
		clui.NewUI,

		wire.FieldsOf(new(*cc.CC), "Config", "Log", "Context", "ErrGroup", "CancelFunc"),
	)
)

func BuildPolicyApp(logOutput logger.Writer, errOutput logger.ErrWriter, configPath config.Path, overrides config.Overrider) (*PolicyApp, func(), error) {
	wire.Build(
		wire.Struct(new(PolicyApp), "*"),
		policyAppSet,
	)
	return &PolicyApp{}, func() {}, nil
}

func BuildTestPolicyApp(logOutput logger.Writer, errOutput logger.ErrWriter, configPath config.Path, overrides config.Overrider) (*PolicyApp, func(), error) {
	wire.Build(
		wire.Struct(new(PolicyApp), "*"),
		policyAppTestSet,
	)
	return &PolicyApp{}, func() {}, nil
}
