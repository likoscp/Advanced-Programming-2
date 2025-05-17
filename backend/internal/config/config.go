package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConfigServer *ConfigServer
	// ConfigDB     *ConfigDB
}

type ConfigServer struct {
	Addr string `mapstructure:"SERVER_ADDR"`
	AuthAddr string `mapstructure:"AUTH_ADDR"`
}

// type ConfigDB struct {
// 	Host     string `mapstructure:"DB_HOST"`
// 	Addr     string `mapstructure:"DB_ADDR"`
// 	User     string `mapstructure:"DB_USER"`
// 	Password string `mapstructure:"DB_PASSWORD"`
// 	Name     string `mapstructure:"DB_NAME"`
// }

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
	return &Config{
		ConfigServer: &configServer,
	}
}
