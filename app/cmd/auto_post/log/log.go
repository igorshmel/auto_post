package log

import (
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"go.uber.org/fx"
)

// newLogger создаёт новый экземпляр Logger
func newLogger() logger.Logger {
	return logger.New(true)
}

// Module ..
var Module = fx.Options(fx.Provide(newLogger))
