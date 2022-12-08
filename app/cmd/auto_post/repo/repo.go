package repo

import (
	"auto_post/app/internal/adapters/repository"
	"auto_post/app/pkg/config"
	"git.fintechru.org/dfa/dfa_lib/logger"
	"go.uber.org/fx"
)

func newRepository(log logger.Logger, cfg config.Config) (*repository.Repository, error) {
	rep, err := repository.NewRepository(
		repository.WithPostgresRepository(cfg, log))
	if err != nil {
		log.Fatal("failed initialize repository with error: %s", err.Error())
	}
	return rep, err
}

// Module ..
var Module = fx.Options(fx.Provide(newRepository))
