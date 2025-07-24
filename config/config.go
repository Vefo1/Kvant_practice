package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	ExternalAPI struct {
		BaseURL string `mapstructure:"base_url"`
		Token   string `mapstructure:"token"`
	} `mapstructure:"external_api"`
	AppAuth struct {
		Token      string `mapstructure:"token"`
		HeaderName string `mapstructure:"header_name"`
	} `mapstructure:"app_auth"`
	Logging struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logging"`
}

// LoadConfig reads configuration from config.yml
func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")      // Look for config in the current directory
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yml")    // Type of config file

	viper.AutomaticEnv() // Read environment variables that match config keys

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
