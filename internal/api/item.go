package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"price/internal/csv"
	"price/internal/model"
	"price/internal/zip"
)

// GetTotal
// @Summary      Получить все товары
// @Description  Метод позволяет получить все товары
// @Tags         Item
// @Accept		 json
// @Produce      json
// @Success      200 {array} model.DataResponse
// @Router       /api/v0/prices/ [get]
func GetTotal(ctx *gin.Context) {
	var response model.DataResponse
	response, err := Storage.GetTotal()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при получении данных из базы",
			"error":   err.Error(),
		})
	}

	csvFilePath := "test_data.csv"
	zipFilePath := "test_data.zip"

	m := csv.Write(csvFilePath, response)
	if m != nil {
		ctx.JSON(http.StatusBadRequest, m)
	}

	m = zip.Dump(csvFilePath, zipFilePath)
	if m != nil {
		ctx.JSON(http.StatusBadRequest, m)
	}

	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=test_data.zip"))
	ctx.Writer.Header().Add("Content-type", fmt.Sprintf("application/zip"))
	ctx.File(zipFilePath)
	return
}

// AddItems
// @Summary      Добавить продукт
// @Description  Метод позволяет записать в базу данных список продуктов
// @Tags         Item
// @Accept		 json
// @Produce      json
// @Param input body model.Items true "Продукты"
// @Success      200
// @Router       /api/v0/prices/ [post]
func AddItems(ctx *gin.Context) {
	var distFilePath = "response.zip"
	//var items model.Items

	reqFile, err := ctx.FormFile(distFilePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при получении параметров запроса",
			"error":   err.Error(),
		})
	}

	// Сохраняем файл
	err = ctx.SaveUploadedFile(reqFile, distFilePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при сохранении файла",
			"error":   err.Error(),
		})
	}

	zip.UnZip(distFilePath, "")

	//if err := ctx.ShouldBindJSON(&items); err != nil {
	//	fmt.Println(err.Error())
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Что-то с запросом", "err": err.Error()})
	//	return
	//}
	//
	//if err := Storage.AddItems(items); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Не удалось записать.", "err": err.Error()})
	//}
}
