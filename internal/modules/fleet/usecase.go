package fleet

import (
	"context"

	fleetmodel "github.com/FauzanAr/fleet-management/internal/modules/fleet/model"
)

type Usecase interface {
	GetFleet(context.Context, fleetmodel.FleetLastLocationRequest) (*fleetmodel.FleetResponse, error)
	GetFleetHistory(context.Context, fleetmodel.FleetHistoryRequest) (*[]fleetmodel.FleetResponse, error)
}