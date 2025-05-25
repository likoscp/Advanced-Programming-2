package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConfigServer *ConfigServer
	ConfigRedis  *ConfigRedis
}

type ConfigServer struct {
	Addr     string `mapstructure:"SERVER_ADDR"`
	AuthAddr string `mapstructure:"AUTH_ADDR"`
}

type ConfigRedis struct {
	ADDR string `mapstructure:"REDIS_ADDR"`
	PORT string `mapstructure:"REDIS_PORT"`
}

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	configServer := ConfigServer{}

	err = viper.Unmarshal(&configServer)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	configRedis := ConfigRedis{}

	err = viper.Unmarshal(&configRedis)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &Config{
		ConfigServer: &configServer,
		ConfigRedis:  &configRedis,
	}
}
