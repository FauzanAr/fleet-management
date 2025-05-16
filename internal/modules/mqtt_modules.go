package modules

import (
	"context"

	"github.com/FauzanAr/fleet-management/internal/config"
	fleethandler "github.com/FauzanAr/fleet-management/internal/modules/fleet/handlers"
	fleetrepository "github.com/FauzanAr/fleet-management/internal/modules/fleet/repositories"
	fleetusecase "github.com/FauzanAr/fleet-management/internal/modules/fleet/usecases"
	"github.com/FauzanAr/fleet-management/internal/mqtt"
)

func (m *Modules) InitMQTT(cfg config.Config) error {
	ctx := context.Background()

	fleetRepo := fleetrepository.NewFleetRepository(m.log, m.db)
	fleetUsecase := fleetusecase.NewFleetUsecase(m.log, fleetRepo)
	mqttHandler := fleethandler.NewFleetMQTTHandler(m.log, fleetUsecase)

	mqttClient := mqtt.NewMQTTHandler(cfg.MQTT.Host, m.log, mqttHandler, cfg)
	mqttClient.SubscribeFleetLocation(cfg.MQTT.FleetTopic)

	m.log.Info(ctx, "MQTT subscriber initialized", nil)
	return nil
}
