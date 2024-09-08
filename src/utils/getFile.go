package utils

import (
	"log"
	"os"
	"strings"
)

func GetFileContent(filePath string) []string {
	body, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	dataTab := strings.Split(string(body), "\n")
	return dataTab
}
