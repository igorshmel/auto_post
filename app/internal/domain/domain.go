package domain

import (
	"auto_post/app/internal/adapters/port"
	logger "auto_post/app/pkg/log"
)

// Domain --
type Domain struct {
	log            logger.Logger
	parseImage     ParseImageFile
	youTubeMachine YouTubeMachine
}

// NewDefaultDomain - инициализация домена
func NewDefaultDomain(log logger.Logger) *Domain {
	log = log.WithMethod("domain")
	return &Domain{log: log}
}

// DomConfiguration --
type DomConfiguration func(dr *Dom) error

// Dom --
type Dom struct {
	ParseImager port.ParseImager
}

// GetParseImager --
func (ths *Dom) GetParseImager() port.ParseImager {
	return ths.ParseImager
}

// NewDom --
func NewDom(configs ...DomConfiguration) (*Dom, error) {
	dom := &Dom{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		err := cfg(dom)
		if err != nil {
			return nil, err
		}
	}

	return dom, nil
}

// WithDefaultDomain --
func WithDefaultDomain(log logger.Logger) DomConfiguration {
	return func(ths *Dom) error {
		dr := NewDefaultDomain(log)
		ths.ParseImager = dr
		return nil
	}
}
