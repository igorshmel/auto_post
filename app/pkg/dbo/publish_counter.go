package dbo

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/types"
	"time"
)

// PublishCounterDBO --
type PublishCounterDBO struct {
	UUID      string
	Date      int64
	Count     int
	Type      types.PublishTypeEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
