package types

// ApplicationInitiatorType - типы заявок
type ApplicationInitiatorType string

const (
	// ApplicationInvestor - Инвестор
	ApplicationInvestor ApplicationInitiatorType = "investor"
	// ApplicationOperator - Оператор
	ApplicationOperator ApplicationInitiatorType = "operator"
)

func (ths ApplicationInitiatorType) String() string {
	return string(ths)
}
