package port

import (
	"context"

	"auto_post/app/pkg/dto"
)

// InitParseImageUseCase _
type InitParseImageUseCase interface {
	Execute(context.Context, *dto.ParseImageReqDTO) error
}

// DownloadImageUseCase _
type DownloadImageUseCase interface {
	Execute(context.Context, *dto.DownloadImageReqDTO) error
}

// BasisUseCase _
type BasisUseCase interface {
	Execute(context.Context) error
}
