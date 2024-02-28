package db

import (
	"fmt"

	"github.com/hoywu/budgetServer/config"
	"github.com/hoywu/budgetServer/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(config config.DatabaseConfig) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.INFO("Init Database Connection Complete")
	return err
}
