package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"os"
	"github.com/Ceramik9/blog_agregator/internal/config"
	"github.com/Ceramik9/blog_agregator/internal/database"
)

func main() {
	
	// create new state struct
	var s state
	
	// load config and add to state
	gatorconfigPath, err := config.GetFilePath(".gatorconfig.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
	}
	gatorconfig, err := config.Read(gatorconfigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
	s.config = &gatorconfig

	// load database and add to state
	db, err := sql.Open("postgres", s.config.DbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	s.db = dbQueries
	
	// initialise new commands list
	cmds := commands {
		commandNames: make(map[string]func(*state, command) error),
	}
	
	// register login handler
	cmds.register("login", handlerLogin)

	// retister register handler
	cmds.register("register", handlerRegister)

	// register reset handler
	cmds.register("reset", handlerReset)

	//register users handler
	cmds.register("users", handlerUsers)
	
	// register agg handler
	cmds.register("agg", handlerAgg)

	// register add feed handler
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	
	// register print feeds handler
	cmds.register("feeds", handlerFeeds)

	// register feed follow handler
	cmds.register("follow", middlewareLoggedIn(handlerFollow))

	// register following handler
	cmds.register("following", middlewareLoggedIn(handlerFollowing))

	// register unfollow feed handler
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	// register unfollow feed

	// parse user command
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
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

