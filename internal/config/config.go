package config

import (
	"context"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"github.com/FauzanAr/fleet-management/internal/pkg/logger"
)

type Config struct {
	AppEnviroment string `env:"APP_ENVIROMENT" envDefault:"debug"`
	AppPort       string `env:"APP_PORT,required"`
	Postgres      Postgres
	MQTT          MQTT
}

type Postgres struct {
	Host         string `env:"POSTGRES_HOST,required"`
	Port         int    `env:"POSTGRES_PORT,required"`
	Password     string `env:"POSTGRES_PASSWORD,required"`
	Username     string `env:"POSTGRES_USERNAME,required"`
	DatabaseName string `env:"POSTGRES_DATABASE_NAME,required"`
}

type MQTT struct {
	Host       string `env:"MQTT_HOST,required"`
	FleetTopic string `env:"MQTT_FLEET_TOPIC,required"`
	ClientId   string `env:"MQTT_CLIENT_ID,required"`
}

func LoadEnv(ctx context.Context, log logger.Logger) (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error(ctx, "Error while load enviroment", err, nil)
		return Config{}, err
	}

	var conf Config
	err = env.Parse(&conf)
	if err != nil {
		log.Error(ctx, "Error while parsing the enviroment", err, nil)
		fmt.Println(": ", err)
		return Config{}, err
	}

	return conf, nil
}
