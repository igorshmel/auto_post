package models

import "git.fintechru.org/dfa/dfa_lib/models/basis"

const paymentOrderRegistryTable = "payment_order_registries"

// PaymentOrderRegistry - представление реестра платёжных поручений в БД
type PaymentOrderRegistry struct {
	basis.BaseModelSoftDelete

	// FileName - наименование итого файла со списком платёжных поручений
	FileName string `gorm:"not null"`

	// Number - номер реестра ?
	Number uint64 `gorm:"not null"`
}

// TableName возвращает наименование таблицы для хранения реестра платёжных поручений
func (r PaymentOrderRegistry) TableName() string {
	return paymentOrderRegistryTable
}
