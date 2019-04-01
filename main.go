package main

import (
	"new-platform-dashboard/api"
	"new-platform-dashboard/config"
	"new-platform-dashboard/db"
	"new-platform-dashboard/middleware"
	"new-platform-dashboard/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	_ "new-platform-dashboard/docs"
)

// @title Dashboard API Documentation
// @version 1.0
// @description This is a documentation of using RESTfull API for dashboard applications. json version : /swagger/doc.json

// @contact.name Partner Dashboard
// @contact.email gabrian@Ainosi.co.id

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	config.Init("dev")
	log.SetLevel(log.DebugLevel)
	db.Init()
	cm := service.InitCache()

	router := gin.Default()
	router.Use(middleware.Auth(cm))

	api.Init(router, cm)
	api.InitWebsocket(router)

	router.Run(":" + config.App.ServerPort)
}
