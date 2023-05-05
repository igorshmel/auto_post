package postgres

import (
	"auto_post/app/pkg/dbo"
	"auto_post/app/pkg/mapping"
)

// GetByActiveStatus --
func (ths *SQLStore) GetByActiveStatus(dbo *dbo.ManagerDBO) error {
	if ths == nil || ths.db == nil {
		return nil
	}
	if dbo == nil {
		return nil
	}
	model := mapping.ManagerDBOtoModel(dbo)
	if err := ths.db.Where("status = ?", "active").
		Order("RAND()").
		Limit(1).
		Find(&model).Error; err != nil {
		return err
	}

	return nil
}
