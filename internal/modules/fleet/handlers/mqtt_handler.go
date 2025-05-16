package fleethandler

import (
	"context"
	"encoding/json"

	"github.com/FauzanAr/fleet-management/internal/modules/fleet"
	fleetmodel "github.com/FauzanAr/fleet-management/internal/modules/fleet/model"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type FleetMQTTHandler struct {
	log logger.Logger
	uc  fleet.Usecase
}

func NewFleetMQTTHandler(log logger.Logger, uc fleet.Usecase) FleetMQTTHandler {
	return FleetMQTTHandler{
		log: log,
		uc:  uc,
	}
}

func (h *FleetMQTTHandler) SubscriberLastLocation(c mqtt.Client, msg mqtt.Message) {
	var payload fleetmodel.FleetInsertRequest
	ctx := context.Background()
	h.log.Info(ctx, "New message incoming", nil)

	err := json.Unmarshal(msg.Payload(), &payload)
	if err != nil {
		h.log.Error(ctx, "Invalid MQTT payload", err, nil)
		return
	}

	if payload.VehicleID == "" {
		h.log.Error(ctx, "Invalid vehicle_id", err, nil)
		return
	}

	err = h.uc.InsertFleet(ctx, payload)
	if err != nil {
		h.log.Error(ctx, "Failed to insert fleet data", err, nil)
	}

	return
}