package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	appConfig *Config
)

// Config represents the application configuration
type Config struct {
	Port                  int    `mapstructure:"APP_PORT"`
	MongoConnectionString string `mapstructure:"MONGO_CONNECTION_STRING"`
}

// GetConfig return the Application configuration
func GetConfig() *Config {
	if appConfig == nil {
		appConfig = &Config{}
		viper.SetDefault("APP_PORT", 8001)
		viper.SetDefault("MONGO_CONNECTION_STRING", "localhost")

		if os.Getenv("ENVIRONMENT") == "DEV" {
			viper.SetConfigName("config")
			viper.SetConfigType("toml")
			viper.AddConfigPath(".")
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatal("error when reading the config file", err)
			}
		} else {
			viper.AutomaticEnv()
		}

		appConfig.Port = viper.GetInt("APP_PORT")
		appConfig.Port = viper.GetInt("MONGO_CONNECTION_STRING")

		err := viper.Unmarshal(appConfig)
		if err != nil {
			log.Fatal("error when unmarshaling config", err)
		}
	}

	return appConfig

}
