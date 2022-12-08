package ddo

import (
	"time"
)

// ReqFileDDO --
type ReqFileDDO struct {
	FileUUID string
	FileURL  string
	Service  string
	Hash     string
	Status   string
}

// ResFileDDO --
type ResFileDDO struct {
	FileUUID  string
	FileURL   string
	Service   string
	Hash      string
	Status    string
	UpdatedAt *time.Time
	CreatedAt time.Time
}
