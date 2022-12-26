package books

import (
	"net/http"

	"github.com/SJamii/book-rest-application/pkg/common/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h handler) GetBooks(ctx *gin.Context) {
	log.Info(" ----  Into the list of books function  ---- ")
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}
