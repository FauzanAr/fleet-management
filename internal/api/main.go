package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/FauzanAr/fleet-management/internal/config"
	"github.com/FauzanAr/fleet-management/internal/modules"
	"github.com/FauzanAr/fleet-management/internal/pkg/databases/psql"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
	"github.com/FauzanAr/fleet-management/internal/pkg/middleware"
	"github.com/FauzanAr/fleet-management/internal/pkg/validator"
	"github.com/FauzanAr/fleet-management/internal/pkg/wrapper"
)

func StartServer() {
	ctx := context.Background()
	log := logger.NewLogger()
	conf, err := config.LoadEnv(ctx, log)
	if err != nil {
		panic("Error while loading enviroment")
	}

	gin.SetMode(conf.AppEnviroment)

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.GinRequestTrace(log))

	validator.InitCustomValidator()

	connection := postgres.NewPostgres(ctx, conf.Postgres, log)
	db, err := connection.Connect()
	if err != nil {
		panic("Database can't connect")
	}

	httpServer := &http.Server{
		Addr:    ":" + conf.AppPort,
		Handler: server,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	go func() {
		<-quit
		log.Info(ctx, "Server is shutting down...", nil)

		ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Error(ctx, "Server forced to shutdown", err, nil)
		}

		log.Sync()
	}()

	server.GET("/", func(c *gin.Context) {
		wrapper.SendSuccessResponse(c, "Server Up and Running", nil, http.StatusOK)
	})

	modules.NewModules(ctx, server, log, db).Init()

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error(ctx, "Unable to start server", err, nil)
	}
}