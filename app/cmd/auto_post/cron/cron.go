package cron

import (
	"errors"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
	"os"
	"time"
)

var (
	// ErrRunnerTasksIsEmpty ...
	ErrRunnerTasksIsEmpty = errors.New("runner tasks is empty")
)

// Runner - worker, периодически выполняющий заданную задачу
type Runner struct {
	*gocron.Scheduler
}

// newCron...
func newCron() (*Runner, error) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}

	return &Runner{Scheduler: gocron.NewScheduler(loc)}, nil
}

// Run запускает worker
func (r *Runner) Run(done <-chan os.Signal) error {
	if len(r.Jobs()) == 0 {
		return ErrRunnerTasksIsEmpty
	}

	r.StartAsync()

	<-done
	return nil
}

// Module ..
var Module = fx.Options(fx.Provide(newCron))
