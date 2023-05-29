package ddo

import (
	status "auto_post/app/pkg/vars/statuses"
	"time"
)

// CreateRecordRequestDDO --
type CreateRecordRequestDDO struct {
	UUID    string
	URL     string
	AuthURL string
	Service string
	Hash    string
	Status  string
}

// CreateRecordResponseDDO --
type CreateRecordResponseDDO struct {
	UUID      string
	URL       string
	AuthURL   string
	Service   string
	Hash      string
	Status    status.RecordStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
