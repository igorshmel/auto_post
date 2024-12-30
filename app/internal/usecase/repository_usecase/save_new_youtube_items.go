package repository_usecase

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
func (ths SaveNewYoutubeItemsUseCase) Execute(ctx context.Context, log logger.Logger, req *dto.SaveNewYoutubeItemsReqDTO) error {
	msg := fmt.Sprintf
	log.Info("REQ: SaveNewYoutubeItems: %v", req)

	// Получаем список элементов
	items := ths.youtubeDomain.GetItems()
	log.Debug("Retrieved %d items from YouTube domain", len(items)) // Уменьшено шумное логирование

	newVideos := make([]dto.VideoInfo, 0) // Фильтруем только новые видео

	for _, item := range items {
		videoDBO := mapping.ConvertDDOItemToDBO(item)
		videoExists, err := ths.extractor.IsVideoIDExists(ctx, &videoDBO)
		if err != nil {
			return lib.ExtErr(
				errs.UnknownError,
				msg("failed to check if video exists (ID: %s): %s", item.VideoId, err.Error()),
				log,
			)
		}

		// Добавляем только те видео, которых нет в базе данных
		if !videoExists {
			newVideos = append(newVideos, dto.VideoInfo{
				Title:   item.Title,
				VideoId: item.VideoId,
			})
		}
	}

	if len(newVideos) == 0 {
		log.Info("No new videos to save")
		return nil
	}

	// Маппинг и множественное сохранение
	videoDBOs := make([]dbo.SaveNewYoutubeItemsDBO, len(newVideos))
	for i, video := range newVideos {
		videoDBOs[i] = *mapping.YoutubeSaveNewYoutubeItemsDtOtoDBO(&video)
	}

	if err := ths.persister.SaveNewYoutubeItemsBatch(ctx, videoDBOs); err != nil {
		return lib.ExtErr(
			errs.UnknownError,
			msg("failed to save new videos batch: %s", err.Error()),
			log,
		)
	}

	log.Info("Successfully saved %d new videos", len(newVideos))
	return nil
}
