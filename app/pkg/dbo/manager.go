package dbo

import (
	status "auto_post/app/pkg/vars/statuses"
	"time"
)

// ManagerDBO --
type ManagerDBO struct {
	UUID      string
	URL       string
	AuthURL   string
	Service   string
	Hash      string
	Status    status.ManagerStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
