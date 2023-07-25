package chanos

import (
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
)

func newChan() <-chan os.Signal {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	return quit
}

// Module ..
var Module = fx.Options(fx.Provide(newChan))
