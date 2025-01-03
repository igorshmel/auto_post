package dbo

import (
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"
)

// YoutubeItemDBO --
type YoutubeItemDBO struct {
	Title     string
	VideoID   string
	Status    status.RecordStatusEnum
	ID        int64
	UUID      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
