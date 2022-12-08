package dbo

import "time"

// Reserve представление блокировки ДС на САУ для работы с репозиторием
type Reserve struct {
	// ID - идентификатор блокировки ДС
	ID uint64

	// UUID - идентификатор(UUID) блокировки ДС
	UUID string

	// Status - статус блокировки
	Status string

	// Amount - величина блокировки в базовых единицах (копейки)
	Amount uint64

	// AccountID - идентификатор САУ
	AccountID uint64

	// CreatedAt - время создания блокировки ДС
	CreatedAt time.Time

	// UpdatedAt - время обновления блокировки ДС
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" блокировки ДС
	DeletedAt *time.Time
}
