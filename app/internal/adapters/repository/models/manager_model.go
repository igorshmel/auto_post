package models

import (
	"auto_post/app/internal/adapters/repository/models/basis"
	"auto_post/app/pkg/vars/constants"
	"auto_post/app/pkg/vars/statuses"
)

// Manager - модель
type Manager struct {
	URL     string                  `gorm:"column:url;not null"`               // URL - ссылка на файл для скачивания
	AuthURL string                  `gorm:"column:auth_url;not null"`          // AuthURL - ссылка на автора
	Service string                  `gorm:"not null"`                          // Название сервиса, откуда будет скачен файл
	Status  status.RecordStatusEnum `gorm:"type:record_status;column:status;"` // Статус состояния файла
	Hash    string                  `gorm:"column:hash"`                       // Хэш файла, для исключения дубликатов изображения
	basis.BaseModel
}

// TableName возвращает имя таблицы
func (ths Manager) TableName() string {
	return constants.ManagerTableName
}
