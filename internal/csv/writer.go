package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"price/internal/model"
	"strconv"
)

func Write(distFilePath string, data model.DataResponse) gin.H {
	// Создаём файл result.csv
	resultCSV, err := os.Create(distFilePath)
	defer resultCSV.Close()
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при создании файла %s", distFilePath),
			"error":   err.Error(),
		}
	}

	// Создаём писателей
	cswWriter := csv.NewWriter(resultCSV)
	defer cswWriter.Flush()

	// Пишем заголовок в файл result.csv
	err = cswWriter.Write([]string{"total_items", "total_categories", "total_price"})
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при записи зоголовка в файл %s", distFilePath),
			"error":   err.Error(),
		}
	}
	var totalItems = strconv.Itoa(data.TotalItems)
	var totalCategories = strconv.Itoa(data.TotalCategories)
	var totalPrice = strconv.FormatFloat(data.TotalPrice, 'f', -1, 64)

	err = cswWriter.Write([]string{totalItems, totalCategories, totalPrice})
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при записи данных в файл %s", distFilePath),
			"error":   err.Error(),
		}
	}

	// фиксируем состояние файла price.csv
	cswWriter.Flush()

	// Закрываем файл price.csv
	err = resultCSV.Close()
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при закрытии файла %s", distFilePath),
			"error":   err.Error(),
		}
	}

	return nil
}
