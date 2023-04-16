package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	envVars *Environments
)

type Environments struct {
	APIPort      string `mapstructure:"API_PORT"`
	AppName      string `mapstructure:"NEW_RELIC_APP_NAME"`
	MongoAddress string `mapstructure:"MONGO_ADDRESS"`
	DBName       string `mapstructure:"DB_NAME"`
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.SetDefault("API_PORT", "8081")
	viper.SetDefault("MONGO_ADDRESS", "")
	viper.SetDefault("DB_NAME", "user-crud")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Errorf("unable find or read configuration file: %w", err)
	}

	if err := viper.Unmarshal(&envVars); err != nil {
		fmt.Errorf("unable to unmarshal configurations from environment: %w", err)
	}

	return envVars
}
