package dbo

import (
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/types"
	"time"
)

// PublishCounterDBO --
type PublishCounterDBO struct {
	UUID      string
	Date      int64                 `gorm:"not null"`                       // Date - дата публикации
	Count     int                   `gorm:"not null"`                       // Количество публикаций в день
	Type      types.PublishTypeEnum `gorm:"type:publish_type;column:type;"` // Тип публикации (название группы в VK например)
	UpdatedAt *time.Time
	CreatedAt time.Time
}
