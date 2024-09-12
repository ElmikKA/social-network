package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func SaveFile(r *http.Request, name string, method string) (string, error) {
	file, fileHeader, err := r.FormFile("avatar")

	if err != nil {
		fmt.Println("Register error decoding json ", err)
		return "", err
	}
	defer file.Close()

	uniquePath := name + method + fileHeader.Filename

	filepath := filepath.Join("./db/assets", uniquePath)

	outFile, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, file); err != nil {
		return "", err
	}

	return filepath, nil
}
