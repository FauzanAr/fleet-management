package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/FauzanAr/fleet-management/internal/config"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
)

type Postgres struct {
	ctx context.Context
	cfg config.Postgres
	log logger.Logger
	db  *gorm.DB
}

func NewPostgres(ctx context.Context, cfg config.Postgres, log logger.Logger) *Postgres {
	return &Postgres{
		ctx: ctx,
		cfg: cfg,
		log: log,
	}
}

func (pg *Postgres) Connect() (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pg.cfg.Host,
		pg.cfg.Port,
		pg.cfg.Username,
		pg.cfg.Password,
		pg.cfg.DatabaseName)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))
	if err != nil {
		pg.log.Error(pg.ctx, "Error connecting to PostgreSQL", err, nil)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		pg.log.Error(pg.ctx, "Error getting PostgreSQL instance", err, nil)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	pg.db = db
	pg.log.Info(pg.ctx, "Success connect to PostgreSQL!", nil)

	return pg, nil
}

func (pg *Postgres) Close() error {
	sqlDB, err := pg.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (pg *Postgres) GetDatabase() *gorm.DB {
	return pg.db
}
