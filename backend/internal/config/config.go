package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ConfigServer *ConfigServer
	ConfigRedis  *ConfigRedis
	ConfigS3 *ConfigS3
}

type ConfigServer struct {
	Addr     string `mapstructure:"SERVER_ADDR"`
	AuthAddr string `mapstructure:"AUTH_ADDR"`
}

type ConfigRedis struct {
	ADDR string `mapstructure:"REDIS_ADDR"`
	PORT string `mapstructure:"REDIS_PORT"`
}

type ConfigS3 struct {
	Endpoint string `mapstructure:"MINIO_ADDR"`
    AccessKey string`mapstructure:"ACCESSKEY"`
    SecretKey string `mapstructure:"SECRETKEY"`
    Bucket string `mapstructure:"BUCKET"`
    UseSSL bool
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

	configS3 := ConfigS3{}

	err = viper.Unmarshal(&configS3)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	configS3.UseSSL = false


	return &Config{
		ConfigServer: &configServer,
		ConfigRedis:  &configRedis,
		ConfigS3: &configS3,
	}
}
