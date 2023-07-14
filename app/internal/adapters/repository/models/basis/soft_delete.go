package basis

import "time"

// BaseModelSoftDelete определяет общие столбцы, которые должны быть во всех структурах базы данных.
// Сущность должна быть мягко удалена
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt *time.Time `gorm:"<-:update"`
}
