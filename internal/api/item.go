package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	myCsv "price/internal/csv"
	"price/internal/model"
	myZip "price/internal/zip"
)

// GetItems
// @Summary      Получить все товары
// @Description  Метод позволяет получить все товары
// @Tags         Item
// @Accept		 json
// @Produce      json
// @Success      200 {array} model.Items
// @Router       /api/v0/prices/ [get]
func GetItems(ctx *gin.Context) {
	var response model.Items
	response, err := Storage.GetItems()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при получении данных из базы",
			"error":   err.Error(),
		})
	}

	csvFilePath := "data.csv"
	zipFilePath := "response.zip"

	m := myCsv.Write(csvFilePath, response)
	if m != nil {
		ctx.JSON(http.StatusBadRequest, m)
	}

	m = myZip.Dump(csvFilePath, zipFilePath)
	if m != nil {
		ctx.JSON(http.StatusBadRequest, m)
	}

	ctx.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=response.zip"))
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
	var reqFileKey = "file"
	var zipFilePath = "source/test_data.zip"

	// Получаем файл
	reqFileHeader, err := ctx.FormFile(reqFileKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при получении тела запроса",
			"error":   err.Error(),
		})
		return
	}

	// Сохраняем архив на сервер (очень сильно надеемся что это архив)
	err = ctx.SaveUploadedFile(reqFileHeader, zipFilePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Ошибка при сохранении полученного файла %s в %s", reqFileHeader.Filename, zipFilePath),
			"error":   err.Error(),
		})
		return
	}

	// Распаковываем архив
	m := myZip.UnZip(zipFilePath, "source")
	if m != nil {
		ctx.JSON(http.StatusBadRequest, m)
		return
	}

	items, m := myCsv.Read("source/test_data.csv")
	if m != nil {
		ctx.JSON(http.StatusBadRequest, m)
		return
	}

	err = Storage.AddItems(items)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при сохранении данных в базу",
			"error":   err.Error(),
		})
		return
	}

	var response model.DataResponse
	response, err = Storage.GetTotal()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ошибка при получении данных из базы",
			"error":   err.Error(),
		})
	}

	// возвращаем ответ
	ctx.JSON(http.StatusOK, response)
	return
}
