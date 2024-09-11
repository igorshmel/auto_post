package repository_usecase

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/lib"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	"github.com/nuttech/bell/v2"
)

// SaveNewYoutubeItemsUseCase --
type SaveNewYoutubeItemsUseCase struct {
	cfg       config.Config
	log       logger.Logger
	bell      *bell.Events
	persister port.Persister
	extractor port.Extractor
}

// NewSaveNewYoutubeItemsUseCase --
func NewSaveNewYoutubeItemsUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
) port.SaveNewYoutubeItemsUseCase {
	return SaveNewYoutubeItemsUseCase{
		cfg:       cfg,
		log:       log,
		bell:      events,
		persister: persister,
		extractor: extractor,
	}
}

// Execute _
func (ths SaveNewYoutubeItemsUseCase) Execute(ctx context.Context, log logger.Logger, req *dto.SaveNewYoutubeItemsReqDTO) error {
	msg := fmt.Sprintf
	log.Info("REQ: SaveNewYoutubeItems: %v", req)

	for _, videoInfo := range req.VideosInfo {
		isVideoExistsDBO := mapping.ConvertVideosInfoDTOtoDBO(videoInfo)
		isVideoIDExists, err := ths.extractor.IsVideoIDExists(ctx, &isVideoExistsDBO)
		if err != nil {
			return lib.ExtErr(errs.UnknownError, msg("failed to get IsVideoIDExists with error: %s", err.Error()), log)
		}
		if !isVideoIDExists {
			saveNewYoutubeItemsDBO := mapping.YoutubeSaveNewYoutubeItemsDtOtoDBO(&videoInfo)
			if err := ths.persister.SaveNewYoutubeItems(ctx, saveNewYoutubeItemsDBO); err != nil {
				return lib.ExtErr(errs.UnknownError, msg("failed to saveNewYoutubeItems with error: %s", err.Error()), log)
			}
		}
	}

	return nil
}
