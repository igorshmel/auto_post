package youtube

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/youtube"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"go.uber.org/fx"
)

func newYouTubeClient(log logger.Logger, cfg config.Config) (*youtube.YouTubeClient, error) {
	rep, err := youtube.NewYouTubeClient(cfg.YouTubeConfig.YouTubeToken)
	if err != nil {
		log.Fatal("failed initialize repository with error: %s", err.Error())
	}
	return rep, err
}

// Module ..
var Module = fx.Options(fx.Provide(newYouTubeClient))
