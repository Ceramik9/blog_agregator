package main

import (
	"github.com/Ceramik9/blog_agregator/internal/config"
	"github.com/Ceramik9/blog_agregator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

