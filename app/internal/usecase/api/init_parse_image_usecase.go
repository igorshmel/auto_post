package api

import (
	"auto_post/app/pkg/errs"
	"context"
	"fmt"

	"auto_post/app/cmd/auto_post/middleware"
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	"auto_post/app/pkg/mapping"
	"git.fintechru.org/dfa/dfa_lib/logger"
)

// InitParseImageUseCase _
type InitParseImageUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	imager    port.ParseImager
}

// NewInitParseImageUseCase _
func NewInitParseImageUseCase(log logger.Logger, persister port.Persister, extractor port.Extractor, filer port.ParseImager) port.InitParseImageUseCase {
	return InitParseImageUseCase{log: log, persister: persister, extractor: extractor, imager: filer}
}

// Execute _
func (ths InitParseImageUseCase) Execute(ctx context.Context, req *dto.ParseImageReqDTO) error {
	msg := fmt.Sprintf
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("usecase InitParseImage")

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	reqParseImageDDO := mapping.ParseImageDTOtoDDO(req)
	resParseImageDDO := ths.imager.InitParseImage(reqParseImageDDO)

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
