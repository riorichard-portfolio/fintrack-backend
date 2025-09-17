package config

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	PGHost           string `validate:"required"`
	PGPort           int    `validate:"required"`
	PGDatabase       string `validate:"required"`
	PGUsername       string `validate:"required"`
	PGPassword       string `validate:"required"`
	PGSSLMode        string `validate:"required"`
	PGChannelBinding string `validate:"required"`
	Port             string `validate:"required"`
}

func mustBeSet(key string) {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("Missing required environment variable: %s", key))
	}
}

func getString(key string) string {
	mustBeSet(key)
	val := viper.GetString(key)
	if val == "" {
		panic(fmt.Sprintf("Missing required environment variable: %s", key))
	}
	return val
}

func getInt(key string) int {
	mustBeSet(key)
	val := viper.GetInt(key)
	if val == 0 {
		panic(fmt.Sprintf("Missing required environment variable: %s", key))
	}
	return val
}

func getBool(key string) bool {
	mustBeSet(key)
	val := viper.GetBool(key)
	return val
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	config := &Config{
		PGHost:           getString("PG_HOST"),
		PGPort:           getInt("PG_PORT"),
		PGDatabase:       getString("PG_DATABASE"),
		PGUsername:       getString("PG_USERNAME"),
		PGPassword:       getString("PG_PASSWORD"),
		PGSSLMode:        getString("PG_SSLMODE"),
		PGChannelBinding: getString("PG_CHANNELBINDING"),
		Port:             getString("APP_PORT"),
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	return config
}

func (c *Config) DBUrl() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s channel_binding=%s",
		c.PGHost, c.PGPort, c.PGUsername, c.PGPassword, c.PGDatabase, c.PGSSLMode, c.PGChannelBinding,
	)
}
