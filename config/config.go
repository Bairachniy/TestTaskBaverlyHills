package config

import (
	"github.com/joho/godotenv"
	"github.com/tkanos/gonfig"
)

type AppConfig struct {
	ListenAddr string `env:"LISTEN_ADDR"`
	DSN        string `env:"DSN"`
}

func GetConfig(configType interface{}) (err error) {
	_ = godotenv.Load(".env")
	return gonfig.GetConf("", configType)
}
