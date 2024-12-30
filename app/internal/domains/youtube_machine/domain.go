package ytmachine

import (
	"github.com/igorshmel/lic_auto_post/app/internal/adapters/port"
	"github.com/igorshmel/lic_auto_post/app/pkg/config"
	logger "github.com/igorshmel/lic_auto_post/app/pkg/log"
)

// DomainConfiguration --
type DomainConfiguration func(dr *Domain) error

// Domain --
type Domain struct {
	YoutubeMachineDomain port.YoutubeMachineDomain
}

// GetYoutubeMachinePorts --
func (ths *Domain) GetYoutubeMachinePorts() port.YoutubeMachineDomain {
	return ths.YoutubeMachineDomain
}

// NewDomain --
func NewDomain(configs ...DomainConfiguration) (*Domain, error) {
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
func WithEntity(log logger.Logger, cfg config.Config) DomainConfiguration {
	return func(ths *Domain) error {
		dr := NewYoutubeMachine(log, cfg)
		ths.YoutubeMachineDomain = dr
		return nil
	}
}
