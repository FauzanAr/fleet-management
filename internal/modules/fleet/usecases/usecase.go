package fleetusecase

import (
	"context"

	"github.com/FauzanAr/fleet-management/internal/modules/fleet"
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

func (fu FleetUsecase) GetFleet(ctx context.Context, payload fleetmodel.FleetLastLocationRequest) error {
	return nil
}

func (fu FleetUsecase) GetFleetHistory(ctx context.Context, payload fleetmodel.FleetHistoryRequest) error {
	return nil
}
