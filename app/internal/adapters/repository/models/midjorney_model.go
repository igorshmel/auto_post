package models

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models/basis"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
)

// Midjorney - модель
type Midjorney struct {
	ImgURL      string                  `gorm:"column:auth_url;not null"`          // AuthURL - ссылка на автора
	Status      status.RecordStatusEnum `gorm:"type:record_status;column:status;"` // Статус состояния файла
	Description string                  `gorm:"column:description;not null"`       // Description - prompt
	Hash        string                  `gorm:"column:hash"`                       // Хэш файла, для исключения дубликатов изображения
	basis.BaseModel
}

// TableName возвращает имя таблицы
func (ths Midjorney) TableName() string {
	return constants.MidjorneyTableName
}
