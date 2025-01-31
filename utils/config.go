package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DB_CONNECTION       string `mapstructure:"DB_CONNECTION"`
	DB_HOST             string `mapstructure:"DB_HOST"`
	DB_PORT             string `mapstructure:"DB_PORT"`
	DB_USER             string `mapstructure:"DB_USER"`
	DB_PASSWORD         string `mapstructure:"DB_PASSWORD"`
	DB_NAME             string `mapstructure:"DB_NAME"`
	DB_SSL_MODE         string `mapstructure:"DB_SSL_MODE"`
	SERVER_HTTPS        string `mapstructure:"SERVER_HTTPS"`
	SERVER_HOST         string `mapstructure:"SERVER_HOST"`
	SERVER_PORT         string `mapstructure:"SERVER_PORT"`
	SERVER_READ_TIMEOUT string `mapstructure:"SERVER_READ_TIMEOUT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env") // Asegura que busque .env
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("⚠️ No se pudo leer la configuración:", err)
		return Config{}, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("⚠️ No se pudo deserializar la configuración:", err)
		return Config{}, err
	}

	return config, nil
}
