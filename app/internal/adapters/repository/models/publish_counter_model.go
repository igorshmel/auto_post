package models

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models/basis"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
)

// PublishCounter - счетчик публикаций по крону
type PublishCounter struct {
	Date  int64                   `gorm:"not null"`                       // Date - дата публикации
	Count int                     `gorm:"not null"`                       // Количество публикаций в день
	Type  status.RecordStatusEnum `gorm:"type:publish_type;column:type;"` // Тип публикации (название группы в VK например)
	basis.BaseModel
}

// TableName возвращает имя таблицы
func (ths PublishCounter) TableName() string {
	return constants.PublishCounterTableName
}
