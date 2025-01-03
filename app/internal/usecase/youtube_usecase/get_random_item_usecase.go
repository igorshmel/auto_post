package youtube_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	transport "github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/youtube"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"github.com/nuttech/bell/v2"
)

// GetRandomItemUseCase --
type GetRandomItemUseCase struct {
	cfg           config.Config
	log           logger.Logger
	bell          *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	youtubeClient *transport.YouTubeClient
	youtubeDomain port.YoutubeMachineDomain
}

// NewGetRandomItemUseCase --
func NewGetRandomItemUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	youtubeClient *transport.YouTubeClient,
	youtubeDomain port.YoutubeMachineDomain,
) port.GetRandomItemUseCase {
	return GetRandomItemUseCase{
		cfg:           cfg,
		log:           log,
		bell:          events,
		persister:     persister,
		extractor:     extractor,
		youtubeClient: youtubeClient,
		youtubeDomain: youtubeDomain,
	}
}

// Execute _
func (ths GetRandomItemUseCase) Execute(ctx context.Context) error {

	// получает уникальный идентификатор запроса
	// ===============================================================================================================
	requestID, ok := ctx.Value(vars.RequestIDKey).(string)
	if !ok {
		return fmt.Errorf("requestID not found in context")
	}

	// настраивает логирование
	// ===============================================================================================================
	msg := fmt.Sprintf
	l := ths.log.WithMethod(requestID + " | usecase GetRandomItem")
	l.Info("start")

	// получает из БД случайную запись о видеоролике в статусе Active
	// ===============================================================================================================
	var randomItem dbo.YoutubeItemDBO

	err := ths.extractor.GetRandomActiveItem(&randomItem)
	if err != nil {
		return lib.ExtErr(
			errs.ExtGetRandomItem,
			msg("failed to fetch random item: %s", err.Error()), l,
		)
	}

	// вносит информацию о случайной записи о видеоролике в домен
	// ===============================================================================================================
	item := mapping.ConvertDBOToYoutubeDDO(randomItem)
	if item != nil {
		ths.youtubeDomain.KeepRandomItem(requestID, item)
	}

	// событие для сценария запроса
	// ===============================================================================================================
	if err := ths.bell.Ring(constants.GetVideo_VK_Event, ctx); err != nil {
		l.Error("unable send event GetVideo_VK_Event with error: %s", err.Error())
	}
	return err
}
