package storage

import (
	"encoding/json"
	"fmt"
	"os"
)


func GetJsonFromString(data []byte) (map[string]interface{}, error) {
	// Check if the data is a valid JSON
	if !json.Valid(data) {
		return nil, fmt.Errorf("Data is not a valid JSON")
	}

	// Unmarshal the data
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling JSON: %v", err)
	}
	return jsonData, nil
}

func GetJsonFromFile(filePath string) (map[string]interface{}, error) {
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("File does not exist: %v", err)
	}

	// Read the file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	// Get the JSON data
	jsonData, err := GetJsonFromString(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

