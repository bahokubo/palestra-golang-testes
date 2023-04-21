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
	MongoAddress string `mapstructure:"MONGO_ADDRESS"`
	DBName       string `mapstructure:"DB_NAME"`
}

func LoadEnvVars() *Environments {
	viper.SetConfigFile(".env")
	viper.GetString("API_PORT")
	viper.SetDefault("MONGO_ADDRESS", "")
	viper.SetDefault("DB_NAME", "user-crud")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("unable find or read configuration file: %w", err)
	}

	if err := viper.Unmarshal(&envVars); err != nil {
		fmt.Println("unable to unmarshal configurations from environment: %w", err)
	}

	return envVars
}
