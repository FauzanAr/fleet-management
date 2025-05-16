package mqtt

import (
	"time"

	"github.com/FauzanAr/fleet-management/internal/config"
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
	cfg    config.Config
}

func NewMQTTHandler(broker string, log logger.Logger, fu fleethandler.FleetMQTTHandler, cfg config.Config) *MQTTHandler {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(cfg.MQTT.FleetTopic).
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

func (h *MQTTHandler) SubscribeFleetLocation(topic string) {
	h.Client.Subscribe(topic, 1, h.fu.SubscriberLastLocation)
}
