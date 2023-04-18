package basis

import "time"

// BaseModel определяет общие столбцы, которые должны содержаться во всех структурах БД, обычно
type BaseModel struct {
	ID        int64      `gorm:"primaryKey;auto_increment;unique"`
	UUID      string     `gorm:"column:uuid;type:uuid;unique;not null"`
	CreatedAt time.Time  `gorm:"<-:create;index"`
	UpdatedAt *time.Time `gorm:"<-:update"`
}
