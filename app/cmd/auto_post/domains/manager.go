package domains

import (
	manager "auto_post/app/internal/domains/manager"
	"auto_post/app/pkg/config"
	logger "auto_post/app/pkg/log"
	"go.uber.org/fx"
)

func newManagerDomain(log logger.Logger, cfg config.Config) (*manager.Domain, error) {
	dom, err := manager.NewDomain(
		manager.WithEntity(log, cfg))
	if err != nil {
		log.Fatal("failed initialize domain with error: %s", err.Error())
	}
	return dom, err
}

// ManagerDomainModule ..
var ManagerDomainModule = fx.Options(fx.Provide(newManagerDomain))
