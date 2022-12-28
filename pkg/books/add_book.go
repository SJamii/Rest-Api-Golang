package books

import (
	"net/http"

	"github.com/SJamii/book-rest-application/pkg/common/models"
	"github.com/SJamii/book-rest-application/pkg/common/response"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h handler) AddBook(ctx *gin.Context) {
	log.Info(" ---  Into Add book function  --- ")

	body := models.Author{}

	if err := ctx.ShouldBindJSON(&body); err != nil {

		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if result := h.DB.Create(&body); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// responseData, err := json.Marshal(response.Response{Status: 200, Data: "", Message: "Successfully Saved into Db"}) //Converting Go struct into json bytes

	// if err != nil {
	// 	log.Error("Error :: ", err)
	// }

	// ctx.Data(http.StatusOK, "application/json", responseData)

	responseData := response.Response{
		Status:  http.StatusOK,
		Data:    body,
		Message: "Data Saved successfully",
	}

	ctx.JSON(http.StatusOK, responseData)

}
