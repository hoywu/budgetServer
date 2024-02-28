package config

import (
	"fmt"

	"github.com/hoywu/budgetServer/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func LoadConfig() (error, AppConfig) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	log.DEBUG("Loading config file...")

	if err := viper.ReadInConfig(); err != nil {
		log.ERROR(fmt.Sprintf("Error reading config file: %s\n", err))
		return err, AppConfig{}
	}

	log.DEBUG("Config file path: " + viper.ConfigFileUsed())

	var dbConfig DatabaseConfig
	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		log.ERROR(fmt.Sprintf("Error reading config file: %s\n", err))
		return err, AppConfig{}
	}

	log.DEBUG("Database config: " + dbConfig.Host + ":" + dbConfig.Port + "/" + dbConfig.DBName)

	return nil, AppConfig{
		Port:     "8080",
		Database: dbConfig,
	}
}
