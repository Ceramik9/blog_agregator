package main

import (
	"github.com/google/uuid"
	"database/sql"
	"context"
	"time"
	"fmt"
	"errors"
	"github.com/Ceramik9/blog_agregator/internal/database"
)

// commands struct
type commands struct {
	commandNames map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	runCommand, ok := c.commandNames[cmd.name]
	if !ok {
		return fmt.Errorf("Command %s does not exist\n", cmd.name)
	}
	err := runCommand(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.commandNames[name] = f
}

// command struct
type command struct {
	name    string
	args    []string
}

func handlerLogin(s *state, cmd command) error {
	
	// check num of args
	if len(cmd.args) != 1 {
		return errors.New("The login command expects a single argument, the username\n")
	}

	// check if username exist
	name := sql.NullString {
		String: cmd.args[0],
		Valid:  true,
	}
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		return fmt.Errorf("User %s does not exist\n", cmd.args[0])
	}

	// login user
	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("The user %s is now logged in\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	
	// check num of args
	if len(cmd.args) != 1 {
		return errors.New("The register command expects a single argument, the name\n")
	}
	
	//create new user
	name := sql.NullString {
		String: cmd.args[0],
		Valid:  true,
	}
	ctx := context.Background()
	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		return fmt.Errorf("User %s already exist\n", cmd.args[0])
	}
	userParams := database.CreateUserParams {
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	_, err = s.db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}
	fmt.Printf("User %s created\n", cmd.args[0])
	
	// update the logged in user in gatorconfig
	err = s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	// success
	return nil
}

