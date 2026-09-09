package config

import (
	"encoding/json"
	"os"
	"fmt"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c Config) SetUser(userName string) error {
	
	// set CurrentUserName
	c.CurrentUserName = userName

	// Marshal JSON
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	
	// set .gatorconfig.json file path
	filePath, err := GetFilePath(".gatorconfig.json")
	if err != nil {
		fmt.Printf("Error: %w\n", err)
	}

	// write data to .gatorconfig.sjon
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

