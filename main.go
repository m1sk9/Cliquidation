package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	pdfFiles, err := filepath.Glob(filepath.Join(currentDir, "*.pdf"))
	if err != nil {
		fmt.Println("Error finding PDF files:", err)
		return
	}

	for _, pdfFile := range pdfFiles {
		dirName := strings.TrimSuffix(filepath.Base(pdfFile), ".pdf")

		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			err := os.Mkdir(dirName, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				continue
			}
		}

		newPath := filepath.Join(dirName, filepath.Base(pdfFile))
		err := os.Rename(pdfFile, newPath)
		if err != nil {
			fmt.Println("Error moving file:", err)
		}
	}
}
