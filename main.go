package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hoywu/budgetServer/config"
	"github.com/hoywu/budgetServer/db"
	"github.com/hoywu/budgetServer/log"
	"github.com/hoywu/budgetServer/model"
	"github.com/hoywu/budgetServer/router"
)

func main() {
	appConfig := loadAppConfig()
	initDBConn(appConfig)
	dbAutoMigrate()
	startServer(appConfig.Server)
}

func loadAppConfig() config.AppConfig {
	err, appConfig := config.LoadConfig()
	if err != nil {
		log.FATAL("Load app config error: " + err.Error())
	}
	return *appConfig
}

func initDBConn(appConfig config.AppConfig) {
	err := db.InitDB(appConfig.Database)
	if err != nil {
		log.FATAL("Init database connection error: " + err.Error())
	}
}

func dbAutoMigrate() {
	db.DB.AutoMigrate(
		&model.User{},
		&model.Token{},
		&model.Category{},
		&model.Record{},
		&model.Budget{},
	)
}

func startServer(serverConfig config.ServerConfig) {
	addr := serverConfig.Host + ":" + serverConfig.Port
	log.INFO("Server started at " + addr)
	gin.SetMode(gin.ReleaseMode)
	r := router.InitRouter()
	r.Run(addr)
}
