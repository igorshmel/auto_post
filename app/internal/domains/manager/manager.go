package parsemachine

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"

	"github.com/google/uuid"
	"github.com/igorshmel/lic_auto_post/app/pkg/ddo"
)

// Manager --
type Manager struct {
	log logger.Logger
	cfg config.Config

	UUID    string
	URL     string
	AuthURL string
	Service string
	Hash    string
	Status  string

	updatedAt *time.Time // Дата изменения записи реестра
	createdAt time.Time  // Дата внесения записи в реестр
}

// NewManager - инициализация домена
func NewManager(log logger.Logger, cfg config.Config) *Manager {
	log = log.WithMethod("managerDomain")
	return &Manager{log: log, cfg: cfg}
}

func (ths *Manager) newRecord() {
	ths.UUID = uuid.New().String()

	ths.recordCreatedAt()
	ths.recordUpdatedAt()

}

func (ths *Manager) readRecord() *ddo.CreateRecordResponseDDO {
	return &ddo.CreateRecordResponseDDO{
		UUID:      ths.UUID,
		URL:       ths.URL,
		AuthURL:   ths.AuthURL,
		Service:   ths.Service,
		Status:    status.RecordStatusEnum(ths.Status),
		Hash:      ths.Hash,
		UpdatedAt: ths.updatedAt,
		CreatedAt: ths.createdAt,
	}
}

// CreateRecord --
func (ths *Manager) CreateRecord(ddo *ddo.CreateRecordRequestDDO) *ddo.CreateRecordResponseDDO {
	activeStatus := status.RecordActiveStatus

	h := sha256.New()
	h.Write([]byte(ddo.URL + ddo.AuthURL))
	hashString := base64.StdEncoding.EncodeToString(h.Sum(nil))

	ths.newRecord()
	ths.URL = ddo.URL
	ths.AuthURL = ddo.AuthURL
	ths.Service = ddo.Service
	ths.Status = activeStatus.Str()
	ths.Hash = hashString

	return ths.readRecord()
}

// ProxyRecord --
func (ths *Manager) ProxyRecord(ddo *ddo.ProxyRecordRequestDDO) *ddo.ProxyRecordResponseDDO {
	activeStatus := status.RecordActiveStatus

	h := sha256.New()
	h.Write([]byte(ddo.URL + ddo.AuthURL))
	hashString := base64.StdEncoding.EncodeToString(h.Sum(nil))

	ths.newRecord()
	ths.URL = ddo.URL
	ths.AuthURL = ddo.AuthURL
	ths.Service = ddo.Service
	ths.Status = activeStatus.Str()
	ths.Hash = hashString

	return nil
}

// recordUpdatedAt --
func (ths *Manager) recordUpdatedAt() {
	t := time.Now()
	ths.updatedAt = &t
}

// recordCreatedAt --
func (ths *Manager) recordCreatedAt() {
	t := time.Now()
	ths.createdAt = t
}
