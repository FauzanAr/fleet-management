package modules

import (
	"context"

	"github.com/gin-gonic/gin"

	fleethandler "github.com/FauzanAr/fleet-management/internal/modules/fleet/handlers"
	fleetrepository "github.com/FauzanAr/fleet-management/internal/modules/fleet/repositories"
	fleetusecase "github.com/FauzanAr/fleet-management/internal/modules/fleet/usecases"
	"github.com/FauzanAr/fleet-management/internal/pkg/databases/psql"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
)

type Modules struct {
	ctx    context.Context
	router *gin.Engine
	log    logger.Logger
	db     *postgres.Postgres
}

func NewModules(ctx context.Context, router *gin.Engine, log logger.Logger, db *postgres.Postgres) *Modules {

	return &Modules{
		ctx:    ctx,
		router: router,
		log:    log,
		db:     db,
	}
}

func (m *Modules) Init() error {
	m.InitFleet()
	return nil
}

func (m *Modules) InitFleet() error {
	fleetRepo := fleetrepository.NewFleetRepository(m.log, m.db)
	fleetUsecase := fleetusecase.NewFleetUsecase(m.log, fleetRepo)
	fleetHandler := fleethandler.NewFleetHandlers(m.log, fleetUsecase)

	group := m.router.Group("/api")
	fleetHandler.FleetRoutes(group)
	return nil
}