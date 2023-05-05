package parsemachine

import (
	"auto_post/app/pkg/config"
	"auto_post/app/pkg/events"
	logger "auto_post/app/pkg/log"
	"auto_post/app/pkg/vars/constants"
	status "auto_post/app/pkg/vars/statuses"
	"crypto/sha256"
	"encoding/base64"
	"github.com/nuttech/bell/v2"
	"time"

	"auto_post/app/pkg/ddo"
	"github.com/google/uuid"
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
		Status:    status.ManagerStatusEnum(ths.Status),
		Hash:      ths.Hash,
		UpdatedAt: ths.updatedAt,
		CreatedAt: ths.createdAt,
	}
}

// CreateRecord --
func (ths *Manager) CreateRecord(ddo *ddo.CreateRecordRequestDDO, bell *bell.Events) *ddo.CreateRecordResponseDDO {
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

	// call event event_name
	if err := bell.Ring(
		constants.DownloadImageEventName,
		events.DownloadImageEvent{
			Link:   ths.URL,
			Output: ths.cfg.DownloadMachine.Path + ths.UUID + ".jpg",
		}); err != nil {
		ths.log.Error("unable send event DownloadImage with error: %s", err.Error())
	}

	return ths.readRecord()
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
