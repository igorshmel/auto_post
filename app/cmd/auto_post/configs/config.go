package configs

import (
	"auto_post/app/pkg/config"
	"go.uber.org/fx"
)

// newConfig returns a new config
func newConfig() (config.Config, error) {
	cfg, err := config.Load()
	return cfg, err
}

// Module --
var Module = fx.Options(fx.Provide(newConfig))
