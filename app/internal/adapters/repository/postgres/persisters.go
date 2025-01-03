package postgres

import (
	"context"
	"errors"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
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

// CreateMidRecord --
func (ths *SQLStore) CreateMidRecord(dbo *dbo.RecordMidDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.RecordMidDBOtoModel(dbo)
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

// SaveNewYoutubeItemsBatch сохраняет массив данных о видеороликах в базу данных.
func (ths *SQLStore) SaveNewYoutubeItemsBatch(ctx context.Context, items []dbo.SaveNewYoutubeItemsDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if len(items) == 0 {
		return errors.New(errs.MsgEmptyInputData)
	}

	// Преобразование DBO объектов в модельные объекты
	dbModels := make([]models.YoutubeItemsModel, len(items))
	for i, item := range items {
		dbModels[i] = *mapping.YoutubeSaveNewYoutubeItemsDBOtoModel(&item)
	}

	// Сохранение всех записей
	if err := ths.db.Table(dbModels[0].TableName()).
		Create(&dbModels).Error; err != nil {
		return err
	}

	return nil
}

// SaveYoutubeNextCursor сохраняет next курсор в базу данных.
func (ths *SQLStore) SaveYoutubeNextCursor(ctx context.Context, dbo *dbo.YoutubeNextCursorDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.YoutubeNextCursorDBOtoModel(dbo)
	if err := ths.db.Table(model.TableName()).
		Create(&model).Error; err != nil {
		return err
	}

	return nil
}

// UpdateYoutubeNextCursor обновляет next курсор в базу данных.
func (ths *SQLStore) UpdateYoutubeNextCursor(ctx context.Context, dbo *dbo.YoutubeNextCursorDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.YoutubeNextCursorDBOtoModel(dbo)
	if err := ths.db.Table(model.TableName()).
		UpdateColumns(&model).Error; err != nil {
		return err
	}

	return nil
}

// UpdateStatusToUsed -- обновляет статус записи с Active на Used
func (ths *SQLStore) UpdateStatusToUsed(recordID uint) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}

	res := ths.db.Model(&models.YoutubeItemsModel{}).
		Where("id = ? AND status = ?", recordID, status.RecordActiveStatus).
		Update("status", status.RecordUsedStatus)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New(errs.MsgNotFound)
	}

	return nil
}
