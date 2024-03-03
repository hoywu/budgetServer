package config

import (
	"fmt"

	"github.com/hoywu/budgetServer/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func LoadConfig() (err error, appConfig *AppConfig) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	log.DEBUG("Loading config file...")

	if err = viper.ReadInConfig(); err != nil {
		log.ERROR(fmt.Sprintf("Error reading config file: %s\n", err))
		return
	}

	log.DEBUG("Config file path: " + viper.ConfigFileUsed())

	if err = viper.Unmarshal(&appConfig); err != nil {
		log.ERROR(fmt.Sprintf("Error reading config file: %s\n", err))
		return
	}

	log.DEBUG("Database config: " +
		appConfig.Database.Host + ":" +
		appConfig.Database.Port + "/" +
		appConfig.Database.DBName)
	return
}
