package api

import (
	"context"
	"fmt"

	"auto_post/app/cmd/auto_post/middleware"
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/dto"
	"auto_post/app/pkg/mapping"
	"git.fintechru.org/dfa/dfa_lib/errs/nominals"
	"git.fintechru.org/dfa/dfa_lib/logger"
)

// DownloadImageUseCase _
type DownloadImageUseCase struct {
	log       logger.Logger
	persister port.Persister
	extractor port.Extractor
	filer     port.Filer
}

// NewDownloadImageUseCase _
func NewDownloadImageUseCase(log logger.Logger, persister port.Persister, extractor port.Extractor, filer port.Filer) port.DownloadImageUseCase {
	return DownloadImageUseCase{log: log, persister: persister, extractor: extractor, filer: filer}
}

// Execute _
func (ths DownloadImageUseCase) Execute(ctx context.Context, req *dto.ReqDownloadImage) error {
	msg := fmt.Sprintf
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("usecase saveFile")

	// -- Бизнес логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	reqFileDDO := mapping.FileDTOtoDDO(req)
	resFileDDO := ths.filer.CreateFile(reqFileDDO)

	// -- Инфраструктурная логика --
	// ---------------------------------------------------------------------------------------------------------------------------
	fileDBO := mapping.FileDDOtoDBO(resFileDDO)

	if err := ths.persister.UnitOfWork(func(tx port.Persister) error { // единица работы, транзакция БД

		if err := ths.persister.SaveNewFile(fileDBO); err != nil {
			return extErr(nominals.UnknownError,
				msg("failed to save file entity (UUID: %s) with error: %s", resFileDDO.FileUUID, err.Error()), log)
		}

		if err := ths.persister.UpdateFileStatus(fileDBO); err != nil {
			return extErr(nominals.UnknownError,
				msg("failed to update file entity (UUID: %s) with error: %s", resFileDDO.FileUUID, err.Error()), log)
		}

		return nil
	}); err != nil {
		return err
	}

	log.Debug("create an file (uuid: %s)", reqFileDDO.FileUUID)

	return nil
}
