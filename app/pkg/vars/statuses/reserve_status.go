package status

// ReserveStatus - Статус блокировки лимита
type ReserveStatus string

const (
	// ReserveHold - заблокировано
	ReserveHold ReserveStatus = "hold"
	// ReserveExecuted - исполнено
	ReserveExecuted ReserveStatus = "executed"
	// ReserveCancel - отменено
	ReserveCancel ReserveStatus = "cancel"
)

func (ths ReserveStatus) String() string {
	return string(ths)
}
