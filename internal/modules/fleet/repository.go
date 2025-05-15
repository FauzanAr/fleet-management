package fleet

import (
	"context"

	fleetentity "github.com/FauzanAr/fleet-management/internal/modules/fleet/entities"
)

type Repository interface {
	GetFleet(context.Context, string) (*fleetentity.Fleet, error)
	GetFleetHistory(context.Context, fleetentity.FleetHistoryQuery) (*[]fleetentity.Fleet, error)
	InsertFleet(context.Context, fleetentity.Fleet) error
}