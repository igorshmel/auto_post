package models

import (
	"git.fintechru.org/dfa/dfa_lib/models/basis"
)

// imageTableName - имя таблицы для хранения файлов изображений для поста
const imageTableName = "files"

// FileModel - модель
type FileModel struct {
	FileURL string `gorm:"column:file_url;not null"` // URL - ссылка на файл для скачивания
	Service string `gorm:"not null"`                 // Название сервиса, откуда будет скачен файл
	Hash    string `gorm:"column:hash"`              // Хэш файла, для исключения дубликатов изображения
	Status  string `gorm:"not null"`                 // Статус состояния файла
	basis.BaseModel
}

// NewFileModel - инициализация модели
//func NewFileModel(
//	fileURL string, service string, hash string, status string, updatedAt *time.Time, createdAt time.Time,
//) FileModel {
//	return FileModel{
//		FileURL: fileURL, Service: service, Hash: hash, Status: status,
//		BaseModel: basis.BaseModel{UpdatedAt: updatedAt, CreatedAt: createdAt},
//	}
//}

// TableName возвращает имя таблицы
func (ths FileModel) TableName() string {
	return imageTableName
}
