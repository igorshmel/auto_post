package models

import "git.fintechru.org/dfa/dfa_lib/models/basis"

const nominalAccountTable = "auto_post"

// NominalAccount - представление номинального счёта в БД
type NominalAccount struct {
	basis.BaseModelSoftDelete

	// AccountDetailID - идентификатор записи с банковскими реквизитами
	AccountDetailID uint64 `gorm:"not null"`

	// AccountDetail - банковская информация о номинальном счёте
	AccountDetail AccountDetail

	// Balance - баланс номинального счёта в базовых единицах (копейки)
	Balance int64

	// Currency - код валюты
	Currency string
}

// TableName возвращает наименование таблицы для хранения данных банковских реквизитов номинального счёта
func (na NominalAccount) TableName() string {
	return nominalAccountTable
}
