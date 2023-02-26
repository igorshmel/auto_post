package postgres

import (
	"errors"

	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/mapping"
	"git.fintechru.org/dfa/dfa_lib/errs/nominals"
)

// SaveNewFile --
func (ths *SQLStore) SaveNewFile(dbo *dbo.ParseImageDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(nominals.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(nominals.MsgEmptyInputData)
	}

	fileModel := mapping.ParseImageDBOtoModel(dbo)
	if err := ths.db.Table(fileModel.TableName()).
		Create(&fileModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateFileStatus --
func (ths *SQLStore) UpdateFileStatus(dbo *dbo.ParseImageDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(nominals.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(nominals.MsgEmptyInputData)
	}

	return ths.db.Model(models.ParseImage{}).
		Where("uuid", dbo.FileUUID).
		Update("status", dbo.Status).
		Error
}
