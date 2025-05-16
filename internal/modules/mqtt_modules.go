package modules

import (
	"context"

	fleethandler "github.com/FauzanAr/fleet-management/internal/modules/fleet/handlers"
	fleetrepository "github.com/FauzanAr/fleet-management/internal/modules/fleet/repositories"
	fleetusecase "github.com/FauzanAr/fleet-management/internal/modules/fleet/usecases"
	"github.com/FauzanAr/fleet-management/internal/mqtt"
)

func (m *Modules) InitMQTT() error {
	ctx := context.Background()

	fleetRepo := fleetrepository.NewFleetRepository(m.log, m.db)
	fleetUsecase := fleetusecase.NewFleetUsecase(m.log, fleetRepo)
	mqttHandler := fleethandler.NewFleetMQTTHandler(m.log, fleetUsecase)

	mqttClient := mqtt.NewMQTTHandler("tcp://mqtt-broker:1883", m.log, mqttHandler)
	mqttClient.SubscribeFleetLocation()

	m.log.Info(ctx, "MQTT subscriber initialized", nil)
	return nil
}