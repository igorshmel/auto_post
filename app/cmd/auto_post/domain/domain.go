package domain

import (
	"auto_post/app/internal/domain"
	"git.fintechru.org/dfa/dfa_lib/logger"
	"go.uber.org/fx"
)

func newDom(log logger.Logger) (*domain.Dom, error) {
	dom, err := domain.NewDom(
		domain.WithDefaultDomain(log))
	if err != nil {
		log.Fatal("failed initialize domain with error: %s", err.Error())
	}
	return dom, err
}

// Module ..
var Module = fx.Options(fx.Provide(newDom))
