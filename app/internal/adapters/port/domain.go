package port

import (
	"auto_post/app/pkg/ddo"
)

// ParseImager interface --
type ParseImager interface {
	InitParseImage(ddo *ddo.ParseImageReqDDO) *ddo.ParseImageResDDO
}
