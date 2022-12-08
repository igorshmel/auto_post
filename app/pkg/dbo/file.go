package dbo

import (
	"time"
)

// FileDBO _
type FileDBO struct {
	FileUUID  string
	FileURL   string
	Service   string
	Hash      string
	Status    string
	UpdatedAt *time.Time
	CreatedAt time.Time
}
