package config

import (
	"os"
	"encoding/json"
)

func Read(file string) (Config, error) {
	
	// read file
	data, err := os.ReadFile(file)
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

