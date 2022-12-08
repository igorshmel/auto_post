package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"time"

	"auto_post/app/pkg/ddo"
	"github.com/google/uuid"
)

// File --
type File struct {
	UUID    string
	URL     string
	Service string
	Hash    string
	Status  string

	updatedAt *time.Time // Дата изменения записи реестра
	createdAt time.Time  // Дата внесения записи в реестр
}

func (ths *Domain) newFile() {
	ths.file.UUID = uuid.New().String()

	ths.fileCreatedAt()
	ths.fileUpdatedAt()

}

func (ths *Domain) readFile() *ddo.ResFileDDO {
	return &ddo.ResFileDDO{
		FileUUID:  ths.file.UUID,
		FileURL:   ths.file.URL,
		Service:   ths.file.Service,
		Status:    ths.file.Status,
		Hash:      ths.file.Hash,
		UpdatedAt: ths.file.updatedAt,
		CreatedAt: ths.file.createdAt,
	}
}

// CreateFile --
func (ths *Domain) CreateFile(fileDDO *ddo.ReqFileDDO) *ddo.ResFileDDO {
	h := sha256.New()
	h.Write([]byte(fileDDO.FileURL))
	hashString := base64.StdEncoding.EncodeToString(h.Sum(nil))
	ths.newFile()
	ths.file.URL = fileDDO.FileURL
	ths.file.Service = fileDDO.Service
	ths.file.Status = "active"
	ths.file.Hash = hashString
	return ths.readFile()
}

// fileUpdatedAt --
func (ths *Domain) fileUpdatedAt() {
	t := time.Now()
	ths.file.updatedAt = &t
}

// fileCreatedAt --
func (ths *Domain) fileCreatedAt() {
	t := time.Now()
	ths.file.createdAt = t
}
