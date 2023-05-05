package parsemachine

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/config"
	logger "auto_post/app/pkg/log"
)

// DomConfiguration --
type DomConfiguration func(dr *Domain) error

// Domain --
type Domain struct {
	ManagerDomain port.ManagerDomain
}

// GetManagerPort --
func (ths *Domain) GetManagerPort() port.ManagerDomain {
	return ths.ManagerDomain
}

// NewDomain --
func NewDomain(configs ...DomConfiguration) (*Domain, error) {
	domain := &Domain{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		err := cfg(domain)
		if err != nil {
			return nil, err
		}
	}

	return domain, nil
}

// WithEntity --
func WithEntity(log logger.Logger, cfg config.Config) DomConfiguration {
	return func(ths *Domain) error {
		dr := NewManager(log, cfg)
		ths.ManagerDomain = dr
		return nil
	}
}
