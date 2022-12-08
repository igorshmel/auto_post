package models

import "git.fintechru.org/dfa/dfa_lib/models/basis"

const transactionTable = "transactions"

// Transaction - представление транзакции в БД
type Transaction struct {
	basis.BaseModelSoftDelete

	// NominalAccountID - идентификатор номинального счёта
	NominalAccountID uint64 `gorm:"not null"`

	// NominalAccount - информация о номинальном счёте
	NominalAccount NominalAccount

	// Type - тип транзакции (в БД тип enum, transactionType)
	Type EnumTransactionType `gorm:"type:enum_transaction_type;column:transaction_type;"`

	// Amount - величина транзакции в базовых единицах (копейки)
	Amount uint64

	// Purpose - обоснование транзакции
	Purpose string
}

// TableName возвращает наименование таблицы для хранения транзакций
func (t Transaction) TableName() string {
	return transactionTable
}
