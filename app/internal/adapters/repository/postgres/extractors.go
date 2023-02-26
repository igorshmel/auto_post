package postgres

import (
	"auto_post/app/pkg/dbo"
)

// GetByUUID  --
func (ths *SQLStore) GetByUUID(dbo *dbo.ParseImageDBO) error {
	if ths == nil || ths.db == nil {
		return nil
	}
	if dbo == nil {
		return nil
	}

	return nil
}
