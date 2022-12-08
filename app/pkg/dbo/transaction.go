package dbo

import "time"

// Transaction представление транзакции для работы с репозиторием
type Transaction struct {
	// ID - идентификатор транзакции
	ID uint64

	// UUID - идентификатор(UUID) транзакции
	UUID string

	// AccountID - идентификатор САУ
	AccountID uint64

	// Type - тип транзакции
	Type string

	// Amount - величина транзакции в базовых единицах (копейки)
	Amount uint64

	// OrderID - идентификатор поручения клиента, которое породило данную транзакцию
	OrderID uint64

	// CreatedAt - время создания транзакции
	CreatedAt time.Time

	// UpdatedAt - время обновления транзакции
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" транзакции
	DeletedAt *time.Time
}
