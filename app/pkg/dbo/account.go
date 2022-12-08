package dbo

import "time"

// Account представление САУ для работы с репозиторием
type Account struct {
	// ID - идентификатор САУ
	ID uint64

	// UUID - идентификатор(UUID) САУ
	UUID string

	// ClientUUID - идентификатор клиента
	ClientUUID string

	// NominalAccountUUID - идентификатор номинального счёта
	NominalAccountUUID string

	// Type - тип счёта
	Type string

	// Number - номер САУ
	Number uint64

	// Balance - баланс САУ в базовых единицах (копейки)
	Balance int64

	// IsBlocked - флаг блокировки САУ
	IsBlocked bool

	// CreatedAt - время создания САУ
	CreatedAt time.Time

	// UpdatedAt - время обновления САУ
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" САУ
	DeletedAt *time.Time
}
