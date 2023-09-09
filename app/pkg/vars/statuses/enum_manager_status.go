package status

import (
	"database/sql/driver"
	"fmt"
)

const recordStatusEnum = "record_status"

// RecordStatusEnum - статус
type RecordStatusEnum string

const (
	// RecordActiveStatus --
	RecordActiveStatus RecordStatusEnum = "active"
	// RecordUsedStatus --
	RecordUsedStatus RecordStatusEnum = "using"
)

// Scan ...
func (e *RecordStatusEnum) Scan(value interface{}) error {
	*e = RecordStatusEnum(value.(string))
	return nil
}

// Value ...
func (e *RecordStatusEnum) Value() (driver.Value, error) {
	return string(*e), nil
}

// Str ...
func (e *RecordStatusEnum) Str() string {
	return string(*e)
}

// MigrationSQL формирует SQL для создания перечисления в БД
func (e *RecordStatusEnum) MigrationSQL() string {
	return fmt.Sprintf(`
		DO $$ 
		BEGIN
			IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = '%[1]s') THEN 
				CREATE TYPE %[1]s AS ENUM ('%[2]s', '%[3]s');
			END IF;
		END$$;`, recordStatusEnum, RecordActiveStatus, RecordUsedStatus)
}
