package dbo

import (
	status "auto_post/app/pkg/vars/statuses"
	"time"
)

// ParseImageDBO _
type ParseImageDBO struct {
	UUID      string
	URL       string
	AuthURL   string
	Service   string
	Hash      string
	Status    status.ParseImageStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
