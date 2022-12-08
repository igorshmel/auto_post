package models

import "git.fintechru.org/dfa/dfa_lib/models/basis"

const clientAccountTable = "client_accounts"

// ClientAccount - представление банковских реквизитов клиента в БД
type ClientAccount struct {
	basis.BaseModelSoftDelete

	// ClientUUID - идентификатор(UUID) клиента (внутренний идентификатор пользователя)
	ClientUUID string `gorm:"type:uuid;index;not null'"`

	// AccountDetailID - идентификатор записи с банковскими реквизитами
	AccountDetailID uint64 `gorm:"index;not null"`

	// AccountDetail - банковская информация о пользовательском счёте
	AccountDetail AccountDetail

	// IsActive - флаг активности счёта
	IsActive bool

	// IsDefault - флаг, является ли счёт счётом по умолчанию
	IsDefault bool
}

// TableName возвращает наименование таблицы для хранения данных банковских реквизитов клиента
func (ca ClientAccount) TableName() string {
	return clientAccountTable
}
