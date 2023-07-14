package repo

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
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
