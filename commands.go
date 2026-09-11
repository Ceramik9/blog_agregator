package main


import (
	"fmt"
	"errors"
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
	args []string
}

func handlerLogin(s *state, cmd command) error {
	
	if len(cmd.args) == 0 {
		return errors.New("the login command expects a single argument, the username\n")
	}
	
	err := s.config.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("The user has been set to %s\n", cmd.args[0])
	return nil
}

