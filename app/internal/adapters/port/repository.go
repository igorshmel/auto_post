package port

import (
	"context"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
)

//go:generate mockgen -source=repository.go -destination=../../mocks/repository.go -package=mocks

// Extractor - объект для извлечения данных из БД
type Extractor interface {
	GetByActiveStatus(*dbo.RecordDBO) error
	GetArtPublishCountByDate(ctx context.Context, counterDBO *dbo.PublishCounterDBO) (uint64, error)
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	UpdateRecordStatus(*dbo.RecordDBO) error
	CreateRecord(*dbo.RecordDBO) error
	SetArtPublishCount(ctx context.Context, countDBO *dbo.PublishCounterDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
