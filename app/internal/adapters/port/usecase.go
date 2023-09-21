package port

import (
	"context"

	"github.com/igorshmel/lic_auto_post/app/pkg/dto"
)

// CreateRecordUseCase --
type CreateRecordUseCase interface {
	Execute(context.Context, *dto.CreateRecordReqDTO) error
}

// ProxyRecordUseCase --
type ProxyRecordUseCase interface {
	Execute(context.Context, *dto.ProxyRecordReqDTO) error
}

// VKWallPostUseCase --
type VKWallPostUseCase interface {
	Execute(context.Context) error
}

// DownloadImageUseCase --
type DownloadImageUseCase interface {
	Execute(context.Context, *dto.DownloadImageReqDTO) error
}

// BasisUseCase --
type BasisUseCase interface {
	Execute(context.Context) error
}
