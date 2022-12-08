package models

import (
	"database/sql/driver"
	"fmt"
)

const enumPaymentOrderStatusName = "enum_payment_order_status"

// EnumPaymentOrderStatus - статус платёжного поручения
type EnumPaymentOrderStatus string

const (
	// CreatedPaymentOrderStatus - поручение создано
	CreatedPaymentOrderStatus EnumPaymentOrderStatus = "created"

	// ProcessingPaymentOrderStatus - поручение в обработке
	ProcessingPaymentOrderStatus EnumPaymentOrderStatus = "processing"

	// SuccessPaymentOrderStatus - поручение успешно выполнено
	SuccessPaymentOrderStatus EnumPaymentOrderStatus = "success"

	// FailurePaymentOrderStatus - поручение не было выполнено
	FailurePaymentOrderStatus EnumPaymentOrderStatus = "failure"
)

// Scan ...
func (e *EnumPaymentOrderStatus) Scan(value interface{}) error {
	*e = EnumPaymentOrderStatus(value.(string))
	return nil
}

// Value ...
func (e *EnumPaymentOrderStatus) Value() (driver.Value, error) {
	return string(*e), nil
}

// MigrationSQL формирует SQL для создания перечисления в БД
func (e *EnumPaymentOrderStatus) MigrationSQL() string {
	return fmt.Sprintf(`
		DO $$ 
		BEGIN
			IF NOT EXISTS(SELECT 1 FROM pg_type WHERE typname = '%[1]s') THEN 
				CREATE TYPE %[1]s AS ENUM ('created', 'processing', 'success', 'failure');
			END IF;
		END$$;`, enumPaymentOrderStatusName)
}
