package port

import "auto_post/app/pkg/dbo"

//go:generate mockgen -source=repository.go -destination=../../mocks/repository.go -package=mocks

// Extractor - объект для извлечения данных из БД
type Extractor interface {
	GetByActiveStatus(*dbo.RecordDBO) error
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	UpdateRecordStatus(*dbo.RecordDBO) error
	CreateRecord(*dbo.RecordDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
