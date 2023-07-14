package domains

import (
	vkMachine "github.com/igorshmel/lic_auto_post/app/internal/domains/vk_machine"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"go.uber.org/fx"
)

func newVkMachineDomain(log logger.Logger, cfg config.Config) (*vkMachine.Domain, error) {
	dom, err := vkMachine.NewDomain(
		vkMachine.WithEntity(log, cfg))
	if err != nil {
		log.Fatal("failed initialize domain with error: %s", err.Error())
	}
	return dom, err
}

// VkMachineDomainModule ...
var VkMachineDomainModule = fx.Options(fx.Provide(newVkMachineDomain))
