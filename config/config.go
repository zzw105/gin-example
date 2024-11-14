package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds all configuration variables.
type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

// LoadConfig loads configuration from environment variables or a config file.
func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return config, nil
}

func GetDSN(config Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
}
