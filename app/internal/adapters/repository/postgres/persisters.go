package postgres

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/errs"
	"auto_post/app/pkg/mapping"
	"errors"
)

// CreateRecord --
func (ths *SQLStore) CreateRecord(dbo *dbo.ManagerDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.ManagerDBOtoModel(dbo)
	if err := ths.db.Table(model.TableName()).
		Create(&model).Error; err != nil {
		return err
	}
	return nil
}

// UpdateRecordStatus --
func (ths *SQLStore) UpdateRecordStatus(dbo *dbo.ManagerDBO) error {
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
