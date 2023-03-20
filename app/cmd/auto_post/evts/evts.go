package evts

import (
	"fmt"
	"github.com/nuttech/bell/v2"
	"go.uber.org/fx"
)

func newBell() (*bell.Events, error) {
	// Use via global state
	eventName := "auto_post"

	// make a new events store
	events := bell.New()

	// add listener on event
	events.Listen(eventName, func(msg bell.Message) {
		fmt.Println(msg)
	})

	// wait until the event completes its work
	events.Wait()
	return events, nil
}

// Module ..
var Module = fx.Options(fx.Provide(newBell))
