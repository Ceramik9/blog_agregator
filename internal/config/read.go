package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

func Read(filePath string) (Config, error) {
	
	// read file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config {}, err
	}

	// unmarshal json
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config {}, err
	}

	return config, nil
}


func GetFilePath(fileName string) (string, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(workingDir, fileName)
	return filePath, nil
}
