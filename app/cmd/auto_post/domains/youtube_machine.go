package domains

import (
	youtubeMachine "github.com/igorshmel/lic_auto_post/app/internal/domains/youtube_machine"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"go.uber.org/fx"
)

func newYoutubeMachineDomain(log logger.Logger, cfg config.Config) (*youtubeMachine.Domain, error) {
	dom, err := youtubeMachine.NewDomain(
		youtubeMachine.WithEntity(log, cfg))
	if err != nil {
		log.Fatal("failed initialize domain with error: %s", err.Error())
	}
	return dom, err
}

// YoutubeMachineDomainModule ...
var YoutubeMachineDomainModule = fx.Options(fx.Provide(newYoutubeMachineDomain))
