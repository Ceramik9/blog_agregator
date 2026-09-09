package main

import (
	"fmt"
	"github.com/Ceramik9/blog_agregator/internal/config"
)

func main() {
	
	// get gatorconfig file path
	configFilePath, err := config.GetFilePath(".gatorconfig.json")
	if err != nil {
		fmt.Printf("Error: %w\n", err)
	}

	// read file
	gatorconfig, err := config.Read(configFilePath)
	if err != nil {
		fmt.Printf("Error: %w\n", err)
	}

	// write file
	gatorconfig.SetUser("sebastian")

	// load .gatorconfig and print content
	gatorconfig, err = config.Read(configFilePath)
	if err != nil {
		fmt.Printf("Error: %w\n", err)
	}
	fmt.Printf("db_url: %s\n", gatorconfig.DbURL)
	fmt.Printf("current_user_name: %s\n", gatorconfig.CurrentUserName)
	fmt.Println("config file struct:")
	fmt.Println(gatorconfig)
}

