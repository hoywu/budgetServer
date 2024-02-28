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
	// load config
	err, appConfig := config.LoadConfig()
	if err != nil {
		log.FATAL(err.Error())
	}

	// init db connection
	err = db.InitDB(appConfig.Database)
	if err != nil {
		log.FATAL(err.Error())
	}

	// create tables
	db.DB.AutoMigrate(&model.User{}, &model.Token{}, &model.Category{}, &model.Record{})

	// start server
	log.INFO("Server started at :8080")
	gin.SetMode(gin.ReleaseMode)
	r := router.InitRouter()
	r.Run(":8080")
}
