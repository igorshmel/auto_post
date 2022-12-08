package api

import (
	"context"

	"auto_post/app/cmd/auto_post/middleware"
	"auto_post/app/internal/adapters/port"
	"git.fintechru.org/dfa/dfa_lib/logger"
)

// BasisUseCase --
type BasisUseCase struct {
	log logger.Logger
}

// NewBasisUseCase _
func NewBasisUseCase(log logger.Logger) port.BasisUseCase {
	return BasisUseCase{log: log}
}

// Execute _
func (ths BasisUseCase) Execute(ctx context.Context) error {
	ths.log = middleware.SetRequestIDPrefix(ctx, ths.log)
	log := ths.log.WithMethod("usecase BasisUsecase")

	log.Debug("success call BasisUsecase")

	return nil
}
