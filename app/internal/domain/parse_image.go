package domain

import (
	status "auto_post/app/pkg/vars/statuses"
	"crypto/sha256"
	"encoding/base64"
	"github.com/nuttech/bell/v2"
	"time"

	"auto_post/app/pkg/ddo"
	"github.com/google/uuid"
)

// ParseImageFile --
type ParseImageFile struct {
	UUID    string
	URL     string
	AuthURL string
	Service string
	Hash    string
	Status  string

	updatedAt *time.Time // Дата изменения записи реестра
	createdAt time.Time  // Дата внесения записи в реестр
}

func (ths *Domain) newParseImage() {
	ths.parseImage.UUID = uuid.New().String()

	ths.parseImageCreatedAt()
	ths.parseImageUpdatedAt()

}

func (ths *Domain) readParseImage() *ddo.ParseImageResDDO {
	return &ddo.ParseImageResDDO{
		UUID:      ths.parseImage.UUID,
		URL:       ths.parseImage.URL,
		AuthURL:   ths.parseImage.AuthURL,
		Service:   ths.parseImage.Service,
		Status:    status.ParseImageStatusEnum(ths.parseImage.Status),
		Hash:      ths.parseImage.Hash,
		UpdatedAt: ths.parseImage.updatedAt,
		CreatedAt: ths.parseImage.createdAt,
	}
}

// InitParseImage --
func (ths *Domain) InitParseImage(ddo *ddo.ParseImageReqDDO, bell *bell.Events) *ddo.ParseImageResDDO {
	activeStatus := status.ParseImageActiveStatus

	h := sha256.New()
	h.Write([]byte(ddo.URL + ddo.AuthURL))
	hashString := base64.StdEncoding.EncodeToString(h.Sum(nil))

	ths.newParseImage()
	ths.parseImage.URL = ddo.URL
	ths.parseImage.AuthURL = ddo.AuthURL
	ths.parseImage.Service = ddo.Service
	ths.parseImage.Status = activeStatus.Str()
	ths.parseImage.Hash = hashString

	// call event event_name
	_ = bell.Ring("auto_post", "Hello bell!")

	return ths.readParseImage()
}

// parseImageUpdatedAt --
func (ths *Domain) parseImageUpdatedAt() {
	t := time.Now()
	ths.parseImage.updatedAt = &t
}

// parseImageCreatedAt --
func (ths *Domain) parseImageCreatedAt() {
	t := time.Now()
	ths.parseImage.createdAt = t
}
