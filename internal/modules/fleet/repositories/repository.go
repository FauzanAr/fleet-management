package fleetrepository

import (
	"context"

	"github.com/FauzanAr/fleet-management/internal/modules/fleet"
	"github.com/FauzanAr/fleet-management/internal/pkg/databases/psql"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
)

type FleetRepository struct {
	db *postgres.Postgres
	log	logger.Logger
}

func NewFleetRepository(log logger.Logger, db *postgres.Postgres) fleet.Repository {
	return FleetRepository{
		db: db,
		log: log,
	}
}

func (r FleetRepository) GetFleet(ctx context.Context) error {
	return nil
}

