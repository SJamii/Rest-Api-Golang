package main

import (
	"github.com/SJamii/book-rest-application/pkg/books"
	"github.com/SJamii/book-rest-application/pkg/common/db"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true

	log.Info("  ---------------   Into the main function  ------------")
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	dbHandler := db.Init(dbUrl)

	books.RegisterRoutes(router, dbHandler)
	router.Run(port)
}
