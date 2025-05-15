package modules

import (
	fleethandler "github.com/FauzanAr/fleet-management/internal/modules/fleet/handlers"
	fleetrepository "github.com/FauzanAr/fleet-management/internal/modules/fleet/repositories"
	fleetusecase "github.com/FauzanAr/fleet-management/internal/modules/fleet/usecases"
	"github.com/FauzanAr/fleet-management/internal/mqtt"
)

func (m *Modules) InitMQTT() error {
	// Init Fleet usecase & MQTT handler
	fleetRepo := fleetrepository.NewFleetRepository(m.log, m.db)
	fleetUsecase := fleetusecase.NewFleetUsecase(m.log, fleetRepo)
	mqttHandler := fleethandler.NewFleetMQTTHandler(m.log, fleetUsecase)

	// Init MQTT client
	mqttClient := mqtt.NewMQTTHandler("tcp://mqtt-broker:1883", m.log, mqttHandler)
	mqttClient.SubscribeFleetLocation()

	m.log.Info(nil, "MQTT subscriber initialized", nil)
	return nil
}