package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"price/internal/model"
	"strconv"
)

func Write(distFilePath string, data model.Items) gin.H {
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
	err = cswWriter.Write([]string{"id", "create_at", "name", "category", "price"})
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при записи зоголовка в файл %s", distFilePath),
			"error":   err.Error(),
		}
	}

	for _, item := range data {
		err = cswWriter.Write([]string{
			strconv.Itoa(item.Id),
			item.CreateDate,
			item.Name,
			item.Category,
			strconv.FormatFloat(item.Price, 'f', -1, 64)})
		if err != nil {
			return gin.H{
				"message": fmt.Sprintf("Ошибка при записи данных в файл %s", distFilePath),
				"error":   err.Error(),
			}
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
