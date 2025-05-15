package fleetrepository

import (
	"context"

	"github.com/FauzanAr/fleet-management/internal/modules/fleet"
	fleetentity "github.com/FauzanAr/fleet-management/internal/modules/fleet/entities"
	postgres "github.com/FauzanAr/fleet-management/internal/pkg/databases/psql"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
	"github.com/FauzanAr/fleet-management/internal/pkg/wrapper"
)

type FleetRepository struct {
	db  *postgres.Postgres
	log logger.Logger
}

func NewFleetRepository(log logger.Logger, db *postgres.Postgres) fleet.Repository {
	return FleetRepository{
		db:  db,
		log: log,
	}
}

func (r FleetRepository) GetFleet(ctx context.Context, vehicleId string) (*fleetentity.Fleet, error) {
	var data fleetentity.Fleet

	err := r.db.
		GetDatabase().
		Where("vehicle_id = ?", vehicleId).
		Order("timestamp ASC").
		First(&data).
		Error

	if err != nil {
		r.log.Error(ctx, "Error getting fleet", err, nil)
		if err.Error() == "record not found" {
			return nil, nil
		}

		return nil, wrapper.InternalServerError("Error getting fleet", nil)
	}

	return &data, nil
}

func (r FleetRepository) GetFleetHistory(ctx context.Context, query fleetentity.FleetHistoryQuery) (*[]fleetentity.Fleet, error) {
	var data []fleetentity.Fleet

	err := r.db.
		GetDatabase().
		Where("vehicle_id = ? AND timestamp BETWEEN ? AND ?", query.VehicleId, query.Start, query.End).
		Order("timestamp ASC").
		Find(&data).
		Error

	if err != nil {
		r.log.Error(ctx, "Error getting fleet history", err, nil)
		return nil, wrapper.InternalServerError("Error getting fleet history", nil)
	}

	return &data, nil
}

func (r FleetRepository) InsertFleet(ctx context.Context, payload fleetentity.Fleet) error {
	err := r.db.
		GetDatabase().
		Create(&payload).
		Error

	if err != nil {
		r.log.Error(ctx, "Error inserting fleet", err, nil)
		return wrapper.InternalServerError("Error inserting fleet", nil)
	}

	return nil
}
