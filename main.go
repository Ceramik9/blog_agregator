package main

import (
	"fmt"
	"github.com/Ceramik9/blog_agregator/internal/config"
)

func main() {
	conf, err := config.Read(".gatorconfig.json")
	if err != nil {
		fmt.Printf("Error: %w\n", err)
	}
	fmt.Printf("db_url: %s\n", conf.DbURL)
}

