package core

import (
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/adapters/persistance"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/services"
)

func Init() (*ports.Drivers, error) {
	repositories, err := persistance.Connect()

	if err != nil {
		return nil, err
	}

	return &ports.Drivers{
		AuthenticationService: services.NewAuthenticationService(),
		FacilityService:       services.NewFacilityService(repositories.FacilityRepository),
	}, nil
}
