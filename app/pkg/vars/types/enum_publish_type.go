package types

import (
	"database/sql/driver"
	"fmt"
)

const publishTypeEnum = "publish_type"

// PublishTypeEnum - тип публикации
type PublishTypeEnum string

const (
	// ArtPublishType --
	ArtPublishType PublishTypeEnum = "art"

	// PhotoshopPublishType --
	PhotoshopPublishType PublishTypeEnum = "photoshop"

	// CodePublishType --
	CodePublishType PublishTypeEnum = "code"
)

// Scan ...
func (e *PublishTypeEnum) Scan(value interface{}) error {
	*e = PublishTypeEnum(value.(string))
	return nil
}

// Value ...
func (e *PublishTypeEnum) Value() (driver.Value, error) {
	return string(*e), nil
}

// Str ...
func (e *PublishTypeEnum) Str() string {
	return string(*e)
}

// MigrationSQL формирует SQL для создания перечисления в БД
func (e *PublishTypeEnum) MigrationSQL() string {
	return fmt.Sprintf(`
		DO $$ 
		BEGIN
			IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = '%[1]s') THEN 
				CREATE TYPE %[1]s AS ENUM ('%[2]s','%[3]s','%[4]s');
			END IF;
		END$$;`, publishTypeEnum, ArtPublishType, PhotoshopPublishType, CodePublishType)
}
