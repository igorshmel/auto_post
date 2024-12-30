package models

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models/basis"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
)

// YoutubeNextCursorModel - модель для таблицы фиксации next курсора
type YoutubeNextCursorModel struct {
	NextPageToken string `json:"nextPageToken,omitempty"`
	basis.BaseModel
}

// TableName возвращает имя таблицы
func (ths YoutubeNextCursorModel) TableName() string {
	return constants.YoutubeNextCursorTableName
}
