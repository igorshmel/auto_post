package api

import (
	"auto_post/app/pkg/config"
	"auto_post/app/pkg/errs"
	"auto_post/app/pkg/events"
	"auto_post/app/pkg/vars/constants"
	"context"
	"fmt"
	"github.com/nuttech/bell/v2"

	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/mapping"
)

// CreateRecordUseCase _
type CreateRecordUseCase struct {
	cfg           config.Config
	log           logger.Logger
	bell          *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	managerDomain port.ManagerDomain
}

// NewCreateRecordUseCase _
func NewCreateRecordUseCase(
	cfg config.Config,
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	managerDomain port.ManagerDomain,
) port.CreateRecordUseCase {
	return CreateRecordUseCase{cfg: cfg, log: log, bell: events, persister: persister, extractor: extractor, managerDomain: managerDomain}
}

// Execute _
func (ths CreateRecordUseCase) Execute(ctx context.Context, req *dto.CreateRecordReqDTO) error {
	msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase CreateRecord")

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	reqCreateRecordDDO := mapping.CreateRecordDTOtoDDO(req)
	resCreateRecordDDO := ths.managerDomain.CreateRecord(reqCreateRecordDDO)

	// -- Инфраструктурная логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	createRecordDBO := mapping.CreateRecordDDOtoDBO(resCreateRecordDDO)
	log.Info("DBO: %v", createRecordDBO)

	if err := ths.persister.UnitOfWork(func(tx port.Persister) error { // единица работы, транзакция БД

		if err := ths.persister.CreateRecord(createRecordDBO); err != nil {
			return extErr(errs.UnknownError,
				msg("failed to create record (UUID: %s) with error: %s", resCreateRecordDDO.UUID, err.Error()), log)
		}

		if err := ths.persister.UpdateRecordStatus(createRecordDBO); err != nil {
			return extErr(errs.UnknownError,
				msg("failed to update record status (UUID: %s) with error: %s", resCreateRecordDDO.UUID, err.Error()), log)
		}

		return nil
	}); err != nil {
		return err
	}

	log.Debug("create record (uuid: %s)", createRecordDBO.UUID)

	// -- Периферия --
	// ---------------------------------------------------------------------------------------------------------------------------

	// отправка события в домен DownloadDomain
	if err := ths.bell.Ring(
		constants.DownloadImageEventName,
		events.DownloadImageEvent{
			Link:   resCreateRecordDDO.URL,
			Output: ths.cfg.DownloadMachine.Path + resCreateRecordDDO.UUID + ".jpg",
		}); err != nil {

		ths.log.Error("unable send event DownloadImage with error: %s", err.Error())
	}

	log.Debug("send DownloadEvent success!")

	return nil
}
