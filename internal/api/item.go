package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"price/internal/model"
)

// GetItems
// @Summary      Получить все товары
// @Description  Метод позволяет получить все товары
// @Tags         Item
// @Accept		 json
// @Produce      json
// @Success      200 {array} model.Items
// @Router       /api/v0/price/ [get]
func GetItems(ctx *gin.Context) {
	var response []string
	response, err := Storage.GetItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Не удалось получить данные.", "err": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"result": response})
}

// AddItems
// @Summary      Добавить продукт
// @Description  Метод позволяет записать в базу данных список продуктов
// @Tags         Item
// @Accept		 json
// @Produce      json
// @Param input body model.Items true "Продукты"
// @Success      200
// @Router       /api/v0/price/ [post]
func AddItems(ctx *gin.Context) {
	var items model.Items
	if err := ctx.ShouldBindJSON(&items); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Что-то с запросом", "err": err.Error()})
		return
	}

	if err := Storage.AddItems(items); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Не удалось записать.", "err": err.Error()})
	}
}
