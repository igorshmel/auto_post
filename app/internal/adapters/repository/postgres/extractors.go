package postgres

import (
	"context"
	"errors"
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/repository/models"
	"github.com/igorshmel/lic_auto_post/app/pkg/dbo"
	"github.com/igorshmel/lic_auto_post/app/pkg/errs"
	"github.com/igorshmel/lic_auto_post/app/pkg/mapping"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/constants"
	status "github.com/igorshmel/lic_auto_post/app/pkg/vars/statuses"
	"github.com/igorshmel/lic_auto_post/app/pkg/vars/types"
	"gorm.io/gorm"
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

// IsVideoIDExists проверяет, существует ли запись с заданным VideoId в базе данных.
func (ths *SQLStore) IsVideoIDExists(ctx context.Context, dbo *dbo.IsVideoExistsDBO) (bool, error) {
	if ths == nil || ths.db == nil {
		return false, errors.New(errs.MsgEmptyDbPointer)
	}
	if dbo == nil {
		return false, errors.New(errs.MsgEmptyInputData)
	}

	model := mapping.ConvertIsVideoExistsDBOtoModel(dbo)
	var count int64
	err := ths.db.Table(model.TableName()).
		Where("video_id = ?", dbo.VideoID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetYoutubeNextCursor - возвращает курсор для списка видеороликов в плейлисте
func (ths *SQLStore) GetYoutubeNextCursor(ctx context.Context) (*dbo.YoutubeNextCursorDBO, error) {
	if ths == nil || ths.db == nil {
		return nil, errors.New(errs.MsgEmptyDbPointer)
	}

	var res dbo.YoutubeNextCursorDBO

	err := ths.db.WithContext(ctx).Table(constants.YoutubeNextCursorTableName).First(&res).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &res, nil
		}
		return nil, err
	}

	return &res, nil
}

// GetRandomActiveItem - получает одну случайную запись со статусом Active
func (ths *SQLStore) GetRandomActiveItem(randomItem *dbo.YoutubeItemDBO) error {
	if ths == nil || ths.db == nil {
		return errors.New(errs.MsgEmptyDbPointer)
	}
	if randomItem == nil {
		return errors.New(errs.MsgEmptyInputData)
	}

	model := models.YoutubeItemsModel{}

	res := ths.db.Model(&model).
		Where("status = ?", status.RecordActiveStatus).
		Order("RANDOM()").
		Limit(1).Find(randomItem)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New(errs.MsgNotFound)
	}

	return nil
}
