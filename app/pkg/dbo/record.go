package dbo

import (
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"
)

// RecordDBO --
type RecordDBO struct {
	UUID      string
	URL       string
	AuthURL   string
	Service   string
	Hash      string
	Status    status.RecordStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
