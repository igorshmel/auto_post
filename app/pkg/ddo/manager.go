package ddo

import (
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
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

// ProxyRecordRequestDDO --
type ProxyRecordRequestDDO struct {
	UUID    string
	URL     string
	AuthURL string
	Service string
	Hash    string
	Status  string
}

// ProxyRecordResponseDDO --
type ProxyRecordResponseDDO struct {
	UUID      string
	URL       string
	AuthURL   string
	Service   string
	Hash      string
	Status    status.RecordStatusEnum
	UpdatedAt *time.Time
	CreatedAt time.Time
}
