package postgres

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/errs"
	"errors"
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
		Where("status = ?", "active").
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
