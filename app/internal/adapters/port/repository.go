package port

import "auto_post/app/pkg/dbo"

//go:generate mockgen -source=repository.go -destination=../../mocks/repository.go -package=mocks

// Extractor - объект для извлечения данных из БД
type Extractor interface {
	GetByUUID(*dbo.ParseImageDBO) error
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	UpdateParseImageStatus(*dbo.ParseImageDBO) error
	InitParseImage(*dbo.ParseImageDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
