package port

import (
	"auto_post/app/pkg/ddo"
	"github.com/nuttech/bell/v2"
)

// ParseImager interface --
type ParseImager interface {
	InitParseImage(ddo *ddo.ParseImageReqDDO, events *bell.Events) *ddo.ParseImageResDDO
}
