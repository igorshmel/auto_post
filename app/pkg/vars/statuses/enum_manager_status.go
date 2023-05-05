package status

import (
	"database/sql/driver"
	"fmt"
)

const managerStatusEnum = "manager_status"

// ManagerStatusEnum - статус
type ManagerStatusEnum string

const (
	// RecordActiveStatus --
	RecordActiveStatus ManagerStatusEnum = "active"
	// ManagerUsingStatus --
	ManagerUsingStatus ManagerStatusEnum = "using"
)

// Scan ...
func (e *ManagerStatusEnum) Scan(value interface{}) error {
	*e = ManagerStatusEnum(value.(string))
	return nil
}

// Value ...
func (e *ManagerStatusEnum) Value() (driver.Value, error) {
	return string(*e), nil
}

// Str ...
func (e *ManagerStatusEnum) Str() string {
	return string(*e)
}

// MigrationSQL формирует SQL для создания перечисления в БД
func (e *ManagerStatusEnum) MigrationSQL() string {
	return fmt.Sprintf(`
		DO $$ 
		BEGIN
			IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = '%[1]s') THEN 
				CREATE TYPE %[1]s AS ENUM ('active', 'using');
			END IF;
		END$$;`, managerStatusEnum)
}
