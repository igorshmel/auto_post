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
	IsVideoIDExists(ctx context.Context, dbo *dbo.IsVideoExistsDBO) (bool, error)
	GetYoutubeNextCursor(ctx context.Context) (*dbo.YoutubeNextCursorDBO, error)
	GetRandomActiveItem(recordDBO *dbo.YoutubeItemDBO) error
}

// Persister - объект для сохранения данных в БД
type Persister interface {
	UpdateRecordStatus(*dbo.RecordDBO) error
	CreateRecord(*dbo.RecordDBO) error
	CreateMidRecord(*dbo.RecordMidDBO) error
	SetArtPublishCount(ctx context.Context, countDBO *dbo.PublishCounterDBO) error
	SaveNewYoutubeItemsBatch(ctx context.Context, dbo []dbo.SaveNewYoutubeItemsDBO) error
	SaveYoutubeNextCursor(ctx context.Context, dbo *dbo.YoutubeNextCursorDBO) error
	UpdateYoutubeNextCursor(ctx context.Context, dbo *dbo.YoutubeNextCursorDBO) error
	UnitOfWork(func(Persister) error) (err error)
}
