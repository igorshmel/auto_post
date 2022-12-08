package dbo

import "time"

// Order представление пользовательского поручения для работы с репозиторием
type Order struct {
	// ID - идентификатор поручения
	ID uint64

	// UUID - идентификатор(UUID) поручения
	UUID string

	// Type - тип поручения
	Type string

	// Status - статус получения
	Status string

	// BasisCode
	BasisCode string

	// Amount - величина перевода в базовых единицах (копейки)
	Amount uint64

	// SenderAccountID - идентификатор САУ отправителя
	SenderAccountID uint64

	// ReserveID - идентификатор блокировки денежных средств под поручение
	ReserveID uint64

	// InitiatorUUID - идентификатор инициатора поручения, в БД тип UUID
	InitiatorUUID string

	// CreatedAt - время создания поручения
	CreatedAt time.Time

	// UpdatedAt - время обновления поручения
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" поручения
	DeletedAt *time.Time
}

// FundingOrder - представление поручения на зачисление ДС для работы с репозиторием
type FundingOrder struct {
	// ID - идентификатор поручения
	ID uint64

	// UUID - идентификатор(UUID) поручения
	UUID string

	// AccountID - идентификатор САУ
	AccountID uint64

	// BankTransactionUUID - идентификатор банковской транзакции
	BankTransactionUUID string

	// NominalAccountUUID - идентификатор номинального счёта
	NominalAccountUUID string

	// CreatedAt - время создания поручения
	CreatedAt time.Time

	// UpdatedAt - время обновления поручения
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" поручения
	DeletedAt *time.Time
}

// PaymentOrder - представление платёжного поручения для работы с репозиторием
type PaymentOrder struct {
	// ID - идентификатор поручения
	ID uint64

	// UUID - идентификатор(UUID) поручения
	UUID string

	// BankClientAccountUUID - идентификатор записи с информацией о клиентском банковском счёте
	BankClientAccountUUID string

	// BankPaymentOrderUUID - идентификатор банковского поручения
	BankPaymentOrderUUID string

	// CreatedAt - время создания поручения
	CreatedAt time.Time

	// UpdatedAt - время обновления поручения
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" поручения
	DeletedAt *time.Time
}

// TransferOrder - представление поручения на перевод между САУ для работы с репозиторием
type TransferOrder struct {
	// ID - идентификатор поручения
	ID uint64

	// UUID - идентификатор(UUID) поручения
	UUID string

	// RecipientAccountID - идентификатор САУ получателя
	RecipientAccountID uint64

	// CreatedAt - время создания поручения
	CreatedAt time.Time

	// UpdatedAt - время обновления поручения
	UpdatedAt *time.Time

	// DeletedAt - время "удаления" поручения
	DeletedAt *time.Time
}
