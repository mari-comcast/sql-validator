package utils

import "os"

func ReadFile(filePath string) (string, error) {
	// Read the file contents
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert the byte slice to string
	fileContent := string(content)

	return fileContent, nil
}
