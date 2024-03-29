package postgres

import (
	"context"
	"errors"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/types"
)

// GetByActiveStatus --
func (ths *SQLStore) GetByActiveStatus(recordDBO *dbo.RecordDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if recordDBO == nil {
		return errors.New(errs.MsgEmptyInputData)
	}
	model := models.Manager{}

	res := ths.db.Model(model).
		Where("status = ?", status.RecordActiveStatus).
		Order("RANDOM()").
		Limit(1).Find(&recordDBO)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(errs.MsgNotFound)
	}

	return nil
}

// GetArtPublishCountByDate --
func (ths *SQLStore) GetArtPublishCountByDate(ctx context.Context, publishCounterDBO *dbo.PublishCounterDBO) (uint64, error) {
	if ths == nil || ths.db == nil {
		return 0, errors.New(errs.MsgEmptyDbPointer)
	}
	if publishCounterDBO == nil {
		return 0, errors.New(errs.MsgEmptyInputData)
	}

	var count uint64
	model := models.PublishCounter{}

	err := ths.db.WithContext(ctx).
		Model(model).
		Where("type = ?", types.ArtPublishType).
		Where("date = ?", publishCounterDBO.Date).
		Select("count").
		Scan(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
