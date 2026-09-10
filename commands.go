package main

import (
	"fmt"
	"errors"
)

type commands struct {
	commandNames map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	// to do
	return nil
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	// to do
}


type command struct {
	name    string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return errors.New("the login handler expects a single argument, the username")
	}
	
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("The user has been set to %s\n", cmd.args[0])
	return nil
}
