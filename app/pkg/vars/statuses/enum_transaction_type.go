package status

import (
	"auto_post/app/pkg/vars/constants"
	"database/sql/driver"
	"fmt"
)

const parseImageStatusEnum = "parse_image_status"

// ParseImageStatusEnum - статус
type ParseImageStatusEnum string

const (
	// ParseImageActiveStatus --
	ParseImageActiveStatus ParseImageStatusEnum = "active"
	// ParseImageUsingStatus --
	ParseImageUsingStatus ParseImageStatusEnum = "using"
)

// Scan ...
func (e *ParseImageStatusEnum) Scan(value interface{}) error {
	*e = ParseImageStatusEnum(value.(string))
	return nil
}

// Value ...
func (e *ParseImageStatusEnum) Value() (driver.Value, error) {
	return string(*e), nil
}

// Str ...
func (e *ParseImageStatusEnum) Str() string {
	return string(*e)
}

// MigrationSQL формирует SQL для создания перечисления в БД
func (e *ParseImageStatusEnum) MigrationSQL() string {
	return fmt.Sprintf(`
		DO $$ 
		BEGIN
			IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = '%[1]s') THEN 
				CREATE TYPE %[1]s AS ENUM ('active', 'using');
			END IF;
		END$$;`, constants.ParseImageTableName)
}
