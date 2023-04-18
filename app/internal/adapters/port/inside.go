package port

import (
	"auto_post/app/pkg/dto"
)

// DownloadImageInside _
type DownloadImageInside interface {
	Execute(*dto.DownloadImageReqDTO) error
}
