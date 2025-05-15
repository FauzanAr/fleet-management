package modules

import (
	"context"

	"github.com/gin-gonic/gin"

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

	return nil
}