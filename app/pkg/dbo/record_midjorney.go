package dbo

import (
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"
)

// RecordMidDBO --
type RecordMidDBO struct {
	ImgURL      string
	Hash        string
	Status      status.RecordStatusEnum
	Description string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
}
