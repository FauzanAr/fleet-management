package mqtt

import (
	"time"

	fleethandler "github.com/FauzanAr/fleet-management/internal/modules/fleet/handlers"
	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type FleetLocationSubscriber interface {
	SubscriberLastLocation(mqtt.Client, mqtt.Message)
}

type MQTTHandler struct {
	Client mqtt.Client
	Logger logger.Logger
	fu     fleethandler.FleetMQTTHandler
}

func NewMQTTHandler(broker string, log logger.Logger, fu fleethandler.FleetMQTTHandler) *MQTTHandler {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID("fleet-tracker-client").
		SetConnectTimeout(10 * time.Second)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Error(nil, "Error connecting to MQTT broker", token.Error(), nil)
		panic("Failed to connect to MQTT broker")
	}

	return &MQTTHandler{
		Client: client,
		Logger: log,
		fu:     fu,
	}
}

func (h *MQTTHandler) SubscribeFleetLocation() {
	topic := "fleet/vehicle/+/location"

	h.Client.Subscribe(topic, 1, h.fu.SubscriberLastLocation)
}
