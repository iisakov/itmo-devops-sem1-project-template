package csv

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"os"
	"price/internal/model"
)

func Read(csvFilePath string) (model.Items, gin.H) {
	var resp model.Items

	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, gin.H{
			"message": fmt.Sprintf("Ошибка при открытии файла %s", csvFilePath),
			"error":   err.Error(),
		}
	}
	defer file.Close()

	err = gocsv.UnmarshalFile(file, &resp)
	if err != nil {
		return nil, gin.H{
			"message": fmt.Sprintf("Ошибка чтении файла %s", csvFilePath),
			"error":   err.Error(),
		}
	}

	return resp, nil
}
