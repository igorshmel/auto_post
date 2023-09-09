package postgres

import (
	"context"
	"errors"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
)

// CreateRecord --
func (ths *SQLStore) CreateRecord(dbo *dbo.RecordDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.RecordDBOtoModel(dbo)
	if err := ths.db.Table(model.TableName()).
		Create(&model).Error; err != nil {
		return err
	}
	return nil
}

// UpdateRecordStatus --
func (ths *SQLStore) UpdateRecordStatus(dbo *dbo.RecordDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	return ths.db.Model(models.Manager{}).
		Where("uuid", dbo.UUID).
		Update("status", dbo.Status).
		Error
}

// SetArtPublishCount --
func (ths *SQLStore) SetArtPublishCount(ctx context.Context, dbo *dbo.PublishCounterDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.SetArtPublishCountDBOtoModel(dbo)
	if err := ths.db.Table(model.TableName()).
		Create(&model).Error; err != nil {
		return err
	}
	return nil
}
