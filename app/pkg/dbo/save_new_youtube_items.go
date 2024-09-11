package dbo

import (
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"
)

// SaveNewYoutubeItemsDBO --
type SaveNewYoutubeItemsDBO struct {
	UUID      string
	Title     string
	VideoId   string
	Status    status.RecordStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
