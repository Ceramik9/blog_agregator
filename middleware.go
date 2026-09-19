package main

import (
	"database/sql"
	"errors"
	"context"
	"github.com/Ceramik9/blog_agregator/internal/database"
)


func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	
	return func(s *state, cmd command) error {
		
		// check if any user is logged in
		if s.config.CurrentUserName == "" {
			return errors.New("user is not logged in")
		}
		// get user
		userName := sql.NullString {
			String: s.config.CurrentUserName,
			Valid:  true,
		}
		user, err := s.db.GetUser(context.Background(), userName)
		if err != nil {
			return err
		}

	return handler(s, cmd, user)
	}
}
