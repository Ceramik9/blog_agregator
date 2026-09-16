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

func  handlerReset(s *state, cmd command) error {
	
	// check num of args
	if len(cmd.args) != 0 {
		return errors.New("The reset command does not take any argumants\n")
	}

	//reset users table
	ctx := context.Background()
	err := s.db.ResetUsers(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Users table has been reset")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	
	// check num of args
	if len(cmd.args) != 0 {
		return errors.New("The users command does not take any argumants\n")
	}
	// get a slice of users
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.String == s.config.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.String)
		} else {
			fmt.Printf("* %s\n",user.String)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {

	// check num of args
	if len(cmd.args) != 0 {
		return errors.New("The users command does not take any argumants\n")
	}
	
	// fatch feed and print content
	url := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	
	//check num of args
	if len(cmd.args) != 2 {
		return errors.New("The addfeed command takes 2 arguments, feed name and the url\n")
	}

	// get user id
	userName := sql.NullString {
		String: s.config.CurrentUserName,
		Valid:  true,
	}
	ctx := context.Background()
	id, err := s.db.GetUserId(ctx, userName)
	if err != nil {
		return err
	}
	userId := uuid.NullUUID {
		UUID:  id,
		Valid: true,
}

	// create feed name
	feedName := sql.NullString {
		String: cmd.args[0],
		Valid:  true,
	}

	// create feed url
	feedUrl := sql.NullString {
		String: cmd.args[1],
		Valid:  true,
	}
	
	// create feed
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    userId,
	}
	_, err = s.db.CreateFeed(ctx, feed)
	if err != nil {
		return err
	}
	fmt.Printf("Feed '%s' created", cmd.args[0])
	return nil
}









