package postgres

import (
	"errors"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
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
