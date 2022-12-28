package books

import (
	"net/http"

	"github.com/SJamii/book-rest-application/pkg/common/models"
	"github.com/SJamii/book-rest-application/pkg/common/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UpdateBookRequestBody struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func (h handler) UpdateBook(ctx *gin.Context) {
	log.Info(" -------------- Update book by Id function  ------------- ")
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		log.Error(" ---- Error Occured during updating  ------- ", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Description = body.Description

	h.DB.Save(&book)

	responseData := response.Response{
		Status:  http.StatusOK,
		Data:    book,
		Message: "Updating Successful",
	}

	ctx.JSON(http.StatusOK, responseData)
}
