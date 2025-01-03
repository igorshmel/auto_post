package youtube_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"github.com/nuttech/bell/v2"
)

// FilteredYoutubeItemsUseCase --
type FilteredYoutubeItemsUseCase struct {
	cfg           config.Config
	log           logger.Logger
	bell          *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	youtubeDomain port.YoutubeMachineDomain
}

// NewFilteredYoutubeItemsUseCase --
func NewFilteredYoutubeItemsUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	youtubeDomain port.YoutubeMachineDomain,
) port.FilteredYoutubeItemsUseCase {
	return FilteredYoutubeItemsUseCase{
		cfg:           cfg,
		log:           log,
		bell:          events,
		persister:     persister,
		extractor:     extractor,
		youtubeDomain: youtubeDomain,
	}
}

// Execute _
func (ths FilteredYoutubeItemsUseCase) Execute(ctx context.Context) error {
	msg := fmt.Sprintf
	requestID, ok := ctx.Value(vars.RequestIDKey).(string)
	if !ok {
		return fmt.Errorf("requestID not found in context")
	}

	l := ths.log.WithMethod(requestID + " | usecase GetPlayList")
	l.Info("start")

	// Получаем список элементов
	items := ths.youtubeDomain.GetItems(requestID)

	// Очищаем список
	ths.youtubeDomain.ClearItems(requestID)

	for _, item := range items {
		videoDBO := mapping.ConvertDDOItemToDBO(item)
		videoExists, err := ths.extractor.IsVideoIDExists(ctx, &videoDBO)
		if err != nil {
			return lib.ExtErr(
				errs.UnknownError,
				msg("failed to check if video exists (ID: %s): %s", item.VideoID, err.Error()),
				l,
			)
		}

		// Добавляем только те видео, которых нет в базе данных
		if !videoExists {
			ths.youtubeDomain.KeepItems(requestID, item.Title, item.VideoID)
		}
	}

	// ~~~~ Периферия ~~~~
	// отправка события Done_Filtered_Items_Event
	if err := ths.bell.Ring(
		constants.Done_Filtered_Items_Event, ctx); err != nil {
		ths.log.Error("unable send event Done_Filtered_Items_Event with error: %s", err.Error())
	}

	return nil
}
