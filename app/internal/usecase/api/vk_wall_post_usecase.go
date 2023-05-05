package api

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/ddo"
	"auto_post/app/pkg/errs"
	logger "auto_post/app/pkg/log"
	"context"
	"fmt"
	"github.com/nuttech/bell/v2"
)

// VKWallPostUseCase --
type VKWallPostUseCase struct {
	log             logger.Logger
	events          *bell.Events
	persister       port.Persister
	extractor       port.Extractor
	vkMachineDomain port.VkMachineDomain
}

// NewVKWallPostUseCase --
func NewVKWallPostUseCase(
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	vkMachineDomain port.VkMachineDomain,
) port.VKWallPostUseCase {
	return VKWallPostUseCase{log: log, events: events, persister: persister, extractor: extractor, vkMachineDomain: vkMachineDomain}
}

// Execute _
func (ths VKWallPostUseCase) Execute(ctx context.Context) error {
	msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase VKMachine")

	parseImageDBO := dbo.ManagerDBO{}
	if err := ths.extractor.GetByActiveStatus(&parseImageDBO); err != nil {
		return extErr(errs.UnknownError,
			msg("failed to get RND entity by active status with error: %s", parseImageDBO.UUID, err.Error()), log)
	}

	vkMachineDDO := ddo.VKMachineDDO{}
	ths.vkMachineDomain.UploadPhotoToServer(&vkMachineDDO)
	return nil
}
