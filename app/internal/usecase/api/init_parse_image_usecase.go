package api

import (
	"auto_post/app/pkg/errs"
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
	log           logger.Logger
	events        *bell.Events
	persister     port.Persister
	extractor     port.Extractor
	managerDomain port.ManagerDomain
}

// NewCreateRecordUseCase _
func NewCreateRecordUseCase(
	log logger.Logger,
	events *bell.Events,
	persister port.Persister,
	extractor port.Extractor,
	managerDomain port.ManagerDomain,
) port.CreateRecordUseCase {
	return CreateRecordUseCase{log: log, events: events, persister: persister, extractor: extractor, managerDomain: managerDomain}
}

// Execute _
func (ths CreateRecordUseCase) Execute(ctx context.Context, req *dto.CreateRecordReqDTO) error {
	msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase CreateRecord")

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	reqCreateRecordDDO := mapping.CreateRecordDTOtoDDO(req)
	resCreateRecordDDO := ths.managerDomain.CreateRecord(reqCreateRecordDDO, ths.events)

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

	return nil
}
