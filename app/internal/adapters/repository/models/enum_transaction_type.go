package models

import (
	"database/sql/driver"
	"fmt"
)

const enumTransactionTypeName = "enum_transaction_type"

// EnumTransactionType - тип транзакции
type EnumTransactionType string

const (
	// DebitTransactionType - транзакция на списание ДС с номинального счёта
	DebitTransactionType EnumTransactionType = "debit"

	// CreditTransactionType - транзакция на пополнение ДС на номинальный счёт
	CreditTransactionType EnumTransactionType = "credit"
)

// Scan ...
func (e *EnumTransactionType) Scan(value interface{}) error {
	*e = EnumTransactionType(value.(string))
	return nil
}

// Value ...
func (e *EnumTransactionType) Value() (driver.Value, error) {
	return string(*e), nil
}

// MigrationSQL формирует SQL для создания перечисления в БД
func (e *EnumTransactionType) MigrationSQL() string {
	return fmt.Sprintf(`
		DO $$ 
		BEGIN
			IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = '%[1]s') THEN 
				CREATE TYPE %[1]s AS ENUM ('debit', 'credit');
			END IF;
		END$$;`, enumTransactionTypeName)
}
