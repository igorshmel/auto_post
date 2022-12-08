package models

import "git.fintechru.org/dfa/dfa_lib/models/basis"

const paymentOrderTable = "payment_orders"

// PaymentOrder - представление платёжного поручения в БД
type PaymentOrder struct {
	basis.BaseModelSoftDelete

	// NominalAccountID - идентификатор номинального счёта
	NominalAccountID uint64 `gorm:"not null"`

	// NominalAccount - информация о номинальном счёте
	NominalAccount NominalAccount

	// PaymentOrderRegistryID - идентификатор записи в реестре платёжных поручений
	PaymentOrderRegistryID uint64

	// PaymentOrderRegistry - информация о реестре платёжных поручений, куда входит текущее поручение
	PaymentOrderRegistry PaymentOrderRegistry

	// Status - статус платёжного поручения
	Status EnumPaymentOrderStatus `gorm:"type:enum_payment_order_status;"`

	// Number - номер платёжного поручения
	Number uint64 `gorm:"not null"`

	// Реквизиты отправителя
	// SenderAccountNumber - номер банковского счёта отправителя
	SenderAccountNumber string `gorm:"not null"`

	// Sender - наименование отправителя
	Sender string `gorm:"not null"`

	// SenderINN - ИНН отправителя
	SenderINN string `gorm:"not null"`

	// SenderKPP - КПП отправителя
	SenderKPP string `gorm:"not null"`

	// SenderBank - наименование банка отправителя
	SenderBank string `gorm:"not null"`

	// SenderBankBIC - БИК банка отправителя
	SenderBankBIC string `gorm:"not null"`

	// SenderBankINN - ИНН банка отправителя
	SenderBankINN string `gorm:"not null"`

	// SenderBankKPP - КПП банка отправителя
	SenderBankKPP string `gorm:"not null"`

	// SenderCorrAccountNumber - корреспондентский банковский счёт отправителя
	SenderCorrAccountNumber string `gorm:"not null"`

	// Реквизиты получателя
	// RecipientAccountNumber - номер банковского счёта получателя
	RecipientAccountNumber string `gorm:"not null"`

	// Recipient - наименование получателя
	Recipient string `gorm:"not null"`

	// RecipientINN - ИНН получателя
	RecipientINN string `gorm:"not null"`

	// RecipientKPP - КПП получателя
	RecipientKPP string `gorm:"not null"`

	// RecipientBank - наименование банка получателя
	RecipientBank string `gorm:"not null"`

	// RecipientBankBIC - БИК банка получателя
	RecipientBankBIC string `gorm:"not null"`

	// RecipientBankINN - ИНН банка получателя
	RecipientBankINN string `gorm:"not null"`

	// RecipientBankKPP - КПП банка получателя
	RecipientBankKPP string `gorm:"not null"`

	// RecipientCorrAccountNumber - корреспондентский банковский счёт получателя
	RecipientCorrAccountNumber string `gorm:"not null"`
}

// TableName возвращает наименование таблицы для хранения платёжных поручений
func (o PaymentOrder) TableName() string {
	return paymentOrderTable
}
