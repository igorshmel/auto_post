package postgres

import (
	"errors"

	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/mapping"
	"git.fintechru.org/dfa/dfa_lib/errs/nominals"
)

// SaveNewFile --
func (ths *SQLStore) SaveNewFile(fileDBO *dbo.FileDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(nominals.MsgEmptyDbPointer)
	}
	if fileDBO == nil {
		return errors.New(nominals.MsgEmptyInputData)
	}

	fileModel := mapping.FileDBOtoModel(fileDBO)
	if err := ths.db.Table(fileModel.TableName()).
		Create(&fileModel).Error; err != nil {
		return err
	}
	return nil
}

// UpdateFileStatus --
func (ths *SQLStore) UpdateFileStatus(fileDBO *dbo.FileDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(nominals.MsgEmptyDbPointer)
	}
	if fileDBO == nil {
		return errors.New(nominals.MsgEmptyInputData)
	}

	return ths.db.Model(models.FileModel{}).
		Where("uuid", fileDBO.FileUUID).
		Update("status", fileDBO.Status).
		Error
}
