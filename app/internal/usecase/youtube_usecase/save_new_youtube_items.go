package youtube_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars"
	"github.com/nuttech/bell/v2"
)

// SaveNewYoutubeItemsUseCase --
type SaveNewYoutubeItemsUseCase struct {
	cfg           config.Config
	log           logger.Logger
	bell          *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	youtubeDomain port.YoutubeMachineDomain
}

// NewSaveNewYoutubeItemsUseCase --
func NewSaveNewYoutubeItemsUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	youtubeDomain port.YoutubeMachineDomain,
) port.SaveNewYoutubeItemsUseCase {
	return SaveNewYoutubeItemsUseCase{
		cfg:           cfg,
		log:           log,
		bell:          events,
		persister:     persister,
		extractor:     extractor,
		youtubeDomain: youtubeDomain,
	}
}

// Execute _
func (ths SaveNewYoutubeItemsUseCase) Execute(ctx context.Context) error {
	msg := fmt.Sprintf
	requestID, ok := ctx.Value(vars.RequestIDKey).(string)
	if !ok {
		return fmt.Errorf("requestID not found in context")
	}

	l := ths.log.WithMethod(requestID + " | usecase GetPlayList")
	l.Info("start")

	// Получаем список элементов
	items := ths.youtubeDomain.GetItems(requestID)

	if len(items) == 0 {
		l.Info("No new videos to save")
		return nil
	}

	// Маппинг и множественное сохранение
	videoDBOs := make([]dbo.SaveNewYoutubeItemsDBO, len(items))
	for i, item := range items {
		videoDBOs[i] = *mapping.YoutubeSaveNewYoutubeItemsDtOtoDBO(&dto.VideoInfo{
			Title:   item.Title,
			VideoID: item.VideoID,
		})
	}

	// Сохраняем новые элементы
	if err := ths.persister.SaveNewYoutubeItemsBatch(ctx, videoDBOs); err != nil {
		return lib.ExtErr(
			errs.UnknownError,
			msg("failed to save new videos batch: %s", err.Error()),
			l,
		)
	}

	ths.youtubeDomain.DeleteState(requestID)

	l.Info("Successfully saved %d new videos")
	return nil
}
