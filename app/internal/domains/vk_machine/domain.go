package vkmachine

import (
	"auto_post/app/internal/adapters/port"
	"auto_post/app/pkg/config"
	logger "auto_post/app/pkg/log"
)

// DomainConfiguration --
type DomainConfiguration func(dr *Domain) error

// Domain --
type Domain struct {
	VkMachineDomain port.VkMachineDomain
}

// GetVkMachinePorts --
func (ths *Domain) GetVkMachinePorts() port.VkMachineDomain {
	return ths.VkMachineDomain
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
		dr := NewVkMachine(log, cfg)
		ths.VkMachineDomain = dr
		return nil
	}
}
