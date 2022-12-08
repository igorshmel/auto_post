package models

import "git.fintechru.org/dfa/dfa_lib/models/basis"

const accountDetailTable = "account_details"

// AccountDetail - представление банковских реквизитов пользователя в БД
type AccountDetail struct {
	basis.BaseModelSoftDelete

	// AccountNumber - номер банковского счёта
	AccountNumber string `gorm:"not null"`

	// Recipient - наименование получателя
	Recipient string `gorm:"not null"`

	// RecipientINN - ИНН получателя
	RecipientINN string `gorm:"not null"`

	// RecipientKPP - КПП получателя
	RecipientKPP string `gorm:"not null"`

	// Bank - наименование банка получателя
	Bank string `gorm:"not null"`

	// BankBIC - БИК банка получателя
	BankBIC string `gorm:"not null"`

	// BankINN - ИНН банка получателя
	BankINN string `gorm:"not null"`

	// BankKPP - КПП банка получателя
	BankKPP string `gorm:"not null"`

	// CorrAccountNumber - корреспондентский банковский счёт
	CorrAccountNumber string `gorm:"not null"`
}

// TableName возвращает наименование таблицы для хранения банковских реквизитов
func (ad AccountDetail) TableName() string {
	return accountDetailTable
}
