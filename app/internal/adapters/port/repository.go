package port

import "auto_post/app/pkg/dbo"

//go:generate mockgen -source=repository.go -destination=../../mocks/repository.go -package=mocks

// Extractor - объект для извлечения данных из БД
type Extractor interface {
	GetByActiveStatus(*dbo.ManagerDBO) error
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	UpdateRecordStatus(*dbo.ManagerDBO) error
	CreateRecord(*dbo.ManagerDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
