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

// InitParseImageUseCase _
type InitParseImageUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	imager    port.ParseImager
	events    *bell.Events
}

// NewInitParseImageUseCase _
func NewInitParseImageUseCase(
	log logger.Logger,
	persister port.Persister,
	extractor port.Extractor,
	filer port.ParseImager,
	events *bell.Events) port.InitParseImageUseCase {
	return InitParseImageUseCase{log: log, persister: persister, extractor: extractor, imager: filer, events: events}
}

// Execute _
func (ths InitParseImageUseCase) Execute(ctx context.Context, req *dto.ParseImageReqDTO) error {
	msg := fmt.Sprintf
	log := ths.log.WithMethod("usecase InitParseImage")

	// call event event_name
	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	reqParseImageDDO := mapping.ParseImageDTOtoDDO(req)
	resParseImageDDO := ths.imager.InitParseImage(reqParseImageDDO, ths.events)

	// -- Инфраструктурная логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	parseImageDBO := mapping.ParseImageDDOtoDBO(resParseImageDDO)
	log.Info("DBO: %v", parseImageDBO)

	if err := ths.persister.UnitOfWork(func(tx port.Persister) error { // единица работы, транзакция БД

		if err := ths.persister.InitParseImage(parseImageDBO); err != nil {
			return extErr(errs.UnknownError,
				msg("failed to save file entity (UUID: %s) with error: %s", resParseImageDDO.UUID, err.Error()), log)
		}

		if err := ths.persister.UpdateParseImageStatus(parseImageDBO); err != nil {
			return extErr(errs.UnknownError,
				msg("failed to update file entity (UUID: %s) with error: %s", resParseImageDDO.UUID, err.Error()), log)
		}

		return nil
	}); err != nil {
		return err
	}

	log.Debug("create an file (uuid: %s)", parseImageDBO.UUID)

	return nil
}
