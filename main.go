package main

import (
	"fmt"
	"os"
	"github.com/Ceramik9/blog_agregator/internal/config"
)

func main() {
	
	// create new state
	gatorconfigPath, err := config.GetFilePath(".gatorconfig.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
	}
	var s state
	gatorconfig, err := config.Read(gatorconfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
	s.config = &gatorconfig
	
	// initialise new commands list
	cmds := commands {
		commandNames: make(map[string]func(*state, command) error),
	}
	
	// register login
	cmds.register("login", handlerLogin)

	// user command
	userCommand := os.Args
	if len(userCommand) < 2 {
		fmt.Fprintln(os.Stderr, "Error: not enough arguments")
		os.Exit(1)
	}
	var cmd command
	cmd.name = userCommand[1]
	cmd.args = userCommand[2:]
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}

