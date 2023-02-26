package ddo

import (
	status "auto_post/app/pkg/vars/statuses"
	"time"
)

// ParseImageReqDDO --
type ParseImageReqDDO struct {
	UUID    string
	URL     string
	AuthURL string
	Service string
	Hash    string
	Status  string
}

// ParseImageResDDO --
type ParseImageResDDO struct {
	UUID      string
	URL       string
	AuthURL   string
	Service   string
	Hash      string
	Status    status.ParseImageStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
