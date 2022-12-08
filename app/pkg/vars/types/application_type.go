package types

// ApplicationType - типы заявок
type ApplicationType string

const (
	// ApplicationInclusion - Включение в реестр
	ApplicationInclusion ApplicationType = "inclusion"
	// ApplicationExclusion - Исключение из реестра
	ApplicationExclusion ApplicationType = "exclusion"
	// ApplicationReport - Выписка из реестра
	ApplicationReport ApplicationType = "report"
)

func (ths ApplicationType) String() string {
	return string(ths)
}
