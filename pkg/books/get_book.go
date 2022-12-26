package books

import (
	"net/http"

	"github.com/SJamii/book-rest-application/pkg/common/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h handler) GetBook(ctx *gin.Context) {
	log.Info(" ----------  Get book by id function ------------- ")
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		log.Error(" -------- Error occured during db call --------- ", result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}
