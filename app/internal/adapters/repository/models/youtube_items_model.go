package models

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models/basis"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
)

// YoutubeItemsModel - список видео роликов
type YoutubeItemsModel struct {
	Title   string                  `gorm:"not null"`
	VideoID string                  `gorm:"not null"`
	Status  status.RecordStatusEnum `gorm:"type:record_status;column:status;"` // статус использования видео ролика
	basis.BaseModel
}

// TableName возвращает имя таблицы
func (ths YoutubeItemsModel) TableName() string {
	return constants.YoutubeItemsTableName
}
