package zip

import (
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnZip(fromFilePath, distPath string) gin.H {

	// Открываем файл fromFilePath
	zipReader, err := zip.OpenReader(fromFilePath)
	if err != nil {
		return gin.H{
			"message": fmt.Sprintf("Ошибка при открытии файла %s", fromFilePath),
			"error":   err.Error(),
		}
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		err = uzf(f, distPath)
		if err != nil {
			return gin.H{
				"message": fmt.Sprintf("Ошибка при распаковке файла %s", fromFilePath),
				"error":   err.Error(),
			}
		}
	}
	return nil
}

func uzf(f *zip.File, distPath string) error {
	filePath := filepath.Join(distPath, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(distPath)+string(os.PathSeparator)) {
		return fmt.Errorf("недопустимый путь к файлу: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err = io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}
