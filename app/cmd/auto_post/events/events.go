package events

import (
	"github.com/nuttech/bell/v2"
	"go.uber.org/fx"
)

func newBell() (*bell.Events, error) {
	events := bell.New()

	// wait until the event completes its work
	events.Wait()
	return events, nil
}

// Module ..
var Module = fx.Options(fx.Provide(newBell))
