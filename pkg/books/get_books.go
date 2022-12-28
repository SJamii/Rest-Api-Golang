package books

import (
	"net/http"

	"github.com/SJamii/book-rest-application/pkg/common/models"
	"github.com/SJamii/book-rest-application/pkg/common/response"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h handler) GetBooks(ctx *gin.Context) {
	log.Info(" ----  Into the list of books function  ---- ")
	var books []models.Book
	result := h.DB.Find(&books)

	if result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// log.Info(" ----------- what --------------", books)
	// b := new(strings.Builder)
	// json.NewEncoder(b).Encode(books)
	// responseData, err := json.Marshal(response.Response{Status: http.StatusOK, Data: books, Message: "Showing all books"}) //Converting Go struct into json bytes

	// if err != nil {
	// 	log.Error("Error :: ", err)
	// }

	responseData := response.Response{
		Status:  http.StatusOK,
		Data:    books,
		Message: "Data Saved successfully",
	}

	ctx.JSON(http.StatusOK, responseData)
	// ctx.Data(http.StatusOK, "application/json", responseData)
}
