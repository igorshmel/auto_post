package log

import (
	"git.fintechru.org/dfa/dfa_lib/logger"
	"go.uber.org/fx"
)

// newLogger создаёт новый экземпляр Logger
func newLogger() logger.Logger {
	return logger.New(true)
}

// Module ..
var Module = fx.Options(fx.Provide(newLogger))
