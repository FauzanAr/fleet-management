package fleetusecase

import (
	"context"

	"github.com/FauzanAr/fleet-management/internal/modules/fleet"
	fleetentity "github.com/FauzanAr/fleet-management/internal/modules/fleet/entities"
	fleethelper "github.com/FauzanAr/fleet-management/internal/modules/fleet/helpers"
	fleetmodel "github.com/FauzanAr/fleet-management/internal/modules/fleet/model"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
)

type FleetUsecase struct {
	fr fleet.Repository
	log logger.Logger
}

func NewFleetUsecase(log logger.Logger, fr fleet.Repository) fleet.Usecase {
	return FleetUsecase{
		fr: fr,
		log: log,
	}
}

func (fu FleetUsecase) GetFleet(ctx context.Context, payload fleetmodel.FleetLastLocationRequest) (*fleetmodel.FleetResponse ,error) {
	var res fleetmodel.FleetResponse

	fleet, err := fu.fr.GetFleet(ctx, payload.VehicleId)
	if err != nil {
		return nil, err
	}

	if fleet == nil {
		return nil, nil
	}

	res = fleethelper.ToFleetResponse(*fleet)

	return &res, nil
}

func (fu FleetUsecase) GetFleetHistory(ctx context.Context, payload fleetmodel.FleetHistoryRequest) (*[]fleetmodel.FleetResponse ,error) {
	var res []fleetmodel.FleetResponse

	fleet, err := fu.fr.GetFleetHistory(ctx, fleetentity.FleetHistoryQuery{
		Start: payload.Start,
		End: payload.End,
		VehicleId: payload.VehicleId,
	})

	if err != nil {
		return nil, err
	}

	if len(*fleet) > 0 {
		res = fleethelper.ToFleetResponseList(*fleet)
	}

	return &res, nil
}

func (fu FleetUsecase) InsertFleet(ctx context.Context, payload fleetmodel.FleetInsertRequest) error {
	var data fleetentity.Fleet

	data.VehicleId = payload.VehicleID
	data.Latitude = payload.Latitude
	data.Longitude = payload.Longitude
	data.Timestamp = payload.Timestamp

	err := fu.fr.InsertFleet(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
