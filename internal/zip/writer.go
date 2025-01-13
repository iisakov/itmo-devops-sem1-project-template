package zip

import (
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Dump(fromFilePath, distFilePath string) gin.H {

	resultCSV, err := os.Open(fromFilePath)
	defer resultCSV.Close()
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при открытии файла %s", fromFilePath),
			"error":   err.Error(),
		}
	}

	// Создаём архива distFilePath
	resultZIP, err := os.Create(distFilePath)
	defer resultZIP.Close()
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при создании архива %s", distFilePath),
			"error":   err.Error(),
		}
	}

	// Создаём писателя в фойл
	zipWriter := zip.NewWriter(resultZIP)
	defer zipWriter.Close()

	// создаём файл price.csv в архиве price.zip
	w, err := zipWriter.Create(fromFilePath)
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при создании пустого файла %s внутри архива %s", fromFilePath, distFilePath),
			"error":   err.Error(),
		}
	}

	// Копируем файл price.csv в архив price.zip/price.csv
	if _, err = io.Copy(w, resultCSV); err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при копировании данных из файла %s в архив %s", fromFilePath, distFilePath),
			"error":   err.Error(),
		}
	}

	return nil
}
