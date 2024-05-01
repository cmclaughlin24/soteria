package core

import (
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/adapters/persistance"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/services"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

func Init() (*ports.Services, error) {
	repositories, err := persistance.Connect()

	if err != nil {
		return nil, err
	}

	return &ports.Services{
		Authentication: services.NewAuthenticationService(&iam.IamHttpClient{AccessTokenUrl: "http://iam/authentication/verify", ApiKeyUrl: "http://iam/api-keys/verify"}),
		Facility:       services.NewFacilityService(repositories.FacilityRepository),
	}, nil
}
