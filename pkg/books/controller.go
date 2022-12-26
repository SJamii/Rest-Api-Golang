package books

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	log.Info("  ----------- Into the routes function  ----------- ")
	h := &handler{
		DB: db,
	}
	routes := router.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/get", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
