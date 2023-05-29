package postgres

import (
	"auto_post/app/internal/adapters/repository/models"
	"auto_post/app/pkg/dbo"
)

// GetByActiveStatus --
func (ths *SQLStore) GetByActiveStatus(recordDBO *dbo.RecordDBO) error {
	if ths == nil || ths.db == nil {
		return nil
	}
	if recordDBO == nil {
		return nil
	}
	model := models.Manager{}
	if err := ths.db.Model(model).
		Where("status = ?", "active").
		Order("RANDOM()").
		Limit(1).Find(&recordDBO).Error; err != nil {
		return err
	}

	return nil
}
