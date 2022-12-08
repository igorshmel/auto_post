package port

import "auto_post/app/pkg/dbo"

//go:generate mockgen -source=repository.go -destination=../../mocks/repository.go -package=mocks

// Extractor - объект для извлечения данных из БД
type Extractor interface {
	GetByUUID(fileDBO *dbo.FileDBO) error
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	UpdateFileStatus(fileDBO *dbo.FileDBO) error
	SaveNewFile(fileDBO *dbo.FileDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
