package port

import (
	"context"

	"auto_post/app/pkg/dto"
)

// DownloadImageUseCase _
type DownloadImageUseCase interface {
	Execute(context.Context, *dto.ReqDownloadImage) error
}

// BasisUseCase _
type BasisUseCase interface {
	Execute(context.Context) error
}
