package inside

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	"auto_post/app/pkg/helpers"
	logger "auto_post/app/pkg/log"
)

// DownloadImageInside _
type DownloadImageInside struct {
	log logger.Logger
}

// NewDownloadImageInside _
func NewDownloadImageInside(
	log logger.Logger,
) port.DownloadImageInside {
	return DownloadImageInside{log: log}
}

// Execute _
func (ths DownloadImageInside) Execute(req *dto.DownloadImageReqDTO) error {
	log := ths.log.WithMethod("inside DownloadImage")

	err := helpers.DownloadFile(req.Output, req.URL)
	if err != nil {
		log.Error("failed to download file with error: %s", err.Error())
	}

	log.Debug("success download file to output folder: %s", req.Output)

	return nil
}
