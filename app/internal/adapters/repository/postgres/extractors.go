package postgres

import (
	"auto_post/app/pkg/dbo"
)

// GetByUUID  --
func (ths *SQLStore) GetByUUID(FileDBO *dbo.FileDBO) error {
	if ths == nil || ths.db == nil {
		return nil
	}
	if FileDBO == nil {
		return nil
	}

	return nil
}
