package api

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	"auto_post/app/pkg/helpers"
	logger "auto_post/app/pkg/log"
	"context"
)

// DownloadImageUseCase _
type DownloadImageUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	imager    port.ParseImager
}

// NewDownloadImageUseCase _
func NewDownloadImageUseCase(
	log logger.Logger,
	persister port.Persister,
	extractor port.Extractor,
	filer port.ParseImager) port.DownloadImageUseCase {
	return DownloadImageUseCase{log: log, persister: persister, extractor: extractor, imager: filer}
}

// Execute _
func (ths DownloadImageUseCase) Execute(ctx context.Context, req *dto.DownloadImageReqDTO) error {
	log := ths.log.WithMethod("usecase DownloadImage")

	err := helpers.DownloadFile(req.Output, req.URL)
	if err != nil {
		log.Error("failed to download file with error: %s", err.Error())
	}

	log.Debug("success download file")

	return nil
}
