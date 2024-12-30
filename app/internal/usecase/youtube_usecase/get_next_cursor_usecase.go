package youtube_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	transport "github.com/igorshmel/lic_auto_post/app/internal/adapters/transport/youtube"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/deo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"github.com/nuttech/bell/v2"
)

// GetNextCursorUseCase --
type GetNextCursorUseCase struct {
	cfg           config.Config
	log           logger.Logger
	bell          *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	youtubeClient *transport.YouTubeClient
	youtubeDomain port.YoutubeMachineDomain
}

// NewGetNextCursorUseCase --
func NewGetNextCursorUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	youtubeClient *transport.YouTubeClient,
	youtubeDomain port.YoutubeMachineDomain,
) port.GetYoutubeNextCursorUseCase {
	return GetNextCursorUseCase{
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
func (ths GetNextCursorUseCase) Execute(ctx context.Context, log logger.Logger) error {
	msg := fmt.Sprintf
	l := ths.log.WithMethod("usecase GetNextCursor")
	l.Info("Try GetGetNextCursor")

	nextCursor, err := ths.extractor.GetYoutubeNextCursor(ctx)
	if err != nil {
		return lib.ExtErr(errs.UnknownError, msg("failed to get nextCursor with error: %s", err.Error()), log)
	}

	if nextCursor != nil {
		fmt.Printf("!!!_ %s", nextCursor.NextPageToken)
	}

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	if nextCursor != nil {
		ths.youtubeDomain.KeepNextCursor(nextCursor.NextPageToken)
	}

	fmt.Printf("%s", ths.youtubeDomain.GetNextCursor())

	// отправка события YouTubeGetPlayListDoneEventName для получения данных по playlist
	if err := ths.bell.Ring(
		constants.YouTubeGetPlayListEventName,
		deo.GetPlayListEvent{
			PlayListID: ths.cfg.YouTubeConfig.YouTubePlayListID,
		}); err != nil {

		ths.log.Error("unable send event DownloadImage with error: %s", err.Error())
	}
	return err
}
