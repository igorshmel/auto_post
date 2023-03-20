package postgres

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/errs"
	"auto_post/app/pkg/mapping"
	"errors"
)

// InitParseImage --
func (ths *SQLStore) InitParseImage(dbo *dbo.ParseImageDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.ParseImageDBOtoModel(dbo)
	if err := ths.db.Table(model.TableName()).
		Create(&model).Error; err != nil {
		return err
	}
	return nil
}

// UpdateParseImageStatus --
func (ths *SQLStore) UpdateParseImageStatus(dbo *dbo.ParseImageDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	return ths.db.Model(models.ParseImage{}).
		Where("uuid", dbo.UUID).
		Update("status", dbo.Status).
		Error
}
