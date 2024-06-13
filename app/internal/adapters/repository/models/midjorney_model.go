package models

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"time"
)

// Midjorney - модель
type Midjorney struct {
	ID          int64                   `gorm:"primaryKey;auto_increment;unique"`
	ImgURL      string                  `gorm:"column:img_url;not null"`           // AuthURL - ссылка на автора
	Status      status.RecordStatusEnum `gorm:"type:record_status;column:status;"` // Статус состояния файла
	Description string                  `gorm:"column:description;not null"`       // Description - prompt
	Hash        string                  `gorm:"column:hash"`                       // Хэш файла, для исключения дубликатов изображения
	CreatedAt   time.Time               `gorm:"<-:create;index"`
	UpdatedAt   *time.Time              `gorm:"<-:update"`
}

// TableName возвращает имя таблицы
func (ths Midjorney) TableName() string {
	return constants.MidjorneyTableName
}
