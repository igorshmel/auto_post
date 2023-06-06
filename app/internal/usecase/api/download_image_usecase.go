package api

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/tools"
	"context"
	"github.com/nuttech/bell/v2"
)

// DownloadImageUseCase _
type DownloadImageUseCase struct {
	log       logger.Logger
	bell      *bell.Events
	persister port.Persister
	extractor port.Extractor
	imager    port.ManagerDomain
}

// NewDownloadImageUseCase _
func NewDownloadImageUseCase(
	log logger.Logger,
	bell *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	filer port.ManagerDomain) port.DownloadImageUseCase {
	return DownloadImageUseCase{log: log, bell: bell, persister: persister, extractor: extractor, imager: filer}
}

// Execute _
func (ths DownloadImageUseCase) Execute(ctx context.Context, req *dto.DownloadImageReqDTO) error {
	log := ths.log.WithMethod("usecase DownloadImage")

	err := tools.DownloadFile(req.Output, req.URL)
	if err != nil {
		log.Error("failed to download file with error: %s", err.Error())
	}

	log.Debug("success download file")

	// -- Периферия --
	// ---------------------------------------------------------------------------------------------------------------------------

	// отправка события в домен VkMachineDomain
	/*	if err := ths.bell.Ring(
			constants.VkWallUploadEventName,
			deo.VkWallUploadEvent{
				FileName: "",
			}); err != nil {

			ths.log.Error("unable send event VkWallUpload with error: %s", err.Error())
		}

		log.Debug("send VkWallUploadEvent success!")*/

	return nil
}
