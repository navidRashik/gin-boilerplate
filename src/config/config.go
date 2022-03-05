package config

import (
	"github.com/spf13/viper"
	"pkg/src/logger"
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// SetupConfig configuration
func SetupConfig() error {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logger.Errorf("Error to reading config file, %s", err)
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Errorf("Unable to decode, %v", err)
		return err
	}

	return nil
}