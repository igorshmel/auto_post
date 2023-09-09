package api

import (
	"context"
	"fmt"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
	"github.com/igorshmel/lic_auto_post/app/pkg/deo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	"github.com/nuttech/bell/v2"

	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
)

// CreateRecordUseCase _
type CreateRecordUseCase struct {
	cfg             config.Config
	log             logger.Logger
	bell            *bell.Events
	persister       port.Persister
	extractor       port.Extractor
	managerDomain   port.ManagerDomain
	vkMachineDomain port.VkMachineDomain
}

// NewCreateRecordUseCase _
func NewCreateRecordUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	managerDomain port.ManagerDomain,
	vkMachineDomain port.VkMachineDomain,
) port.CreateRecordUseCase {
	return CreateRecordUseCase{
		cfg:             cfg,
		log:             log,
		bell:            events,
		persister:       persister,
		extractor:       extractor,
		managerDomain:   managerDomain,
		vkMachineDomain: vkMachineDomain,
	}
}

// Execute _
func (ths CreateRecordUseCase) Execute(ctx context.Context, req *dto.CreateRecordReqDTO) error {
	msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase CreateRecord")

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	reqCreateRecordDDO := mapping.CreateRecordDTOtoDDO(req)
	resCreateRecordDDO := ths.managerDomain.CreateRecord(reqCreateRecordDDO)

	reqGetPath := ddo.VKMachine{FileName: resCreateRecordDDO.UUID}
	path := ths.vkMachineDomain.GetPath(&reqGetPath)

	// -- Инфраструктурная логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	createRecordDBO := mapping.CreateRecordDDOtoDBO(resCreateRecordDDO)
	log.Info("DBO: %v", createRecordDBO)

	if err := ths.persister.CreateRecord(createRecordDBO); err != nil {
		return extErr(errs.UnknownError,
			msg("failed to create record (UUID: %s) with error: %s", resCreateRecordDDO.UUID, err.Error()), log)
	}

	log.Debug("create record (uuid: %s)", createRecordDBO.UUID)

	// -- Периферия --
	// ---------------------------------------------------------------------------------------------------------------------------

	// отправка события в домен DownloadDomain
	if err := ths.bell.Ring(
		constants.DownloadImageEventName,
		deo.DownloadImageEvent{
			Link:   resCreateRecordDDO.URL,
			Output: path,
		}); err != nil {

		ths.log.Error("unable send event DownloadImage with error: %s", err.Error())
	}

	log.Debug("send DownloadEvent success!")

	return nil
}
