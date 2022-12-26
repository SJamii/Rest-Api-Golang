package books

import (
	"net/http"

	"github.com/SJamii/book-rest-application/pkg/common/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h handler) DeleteBook(ctx *gin.Context) {
	log.Info(" ------------- Delete by book id function  --------------")
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		log.Error(" ----- Error Occured while deleting ----------- ", result.Error)
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusOK)
}
