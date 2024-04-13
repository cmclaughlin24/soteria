package persistance

import (
	"fmt"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/persistance/postgres"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
)

func Connect() (*ports.Repositories, error) {
	persistanceType := os.Getenv("PERSISTANCE_TYPE")

	switch persistanceType {
	case "postgres":
		return postgres.Connect()
	default:
		return nil, fmt.Errorf("invalid persistance type: %s is not implemented", persistanceType)
	}
}
