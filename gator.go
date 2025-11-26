package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joshckidd/gator/internal/config"
	"github.com/joshckidd/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No arguments given for login command.")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	s.cfg.CurrentUserName = cmd.args[0]
	err = s.cfg.SetUser()
	if err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("No arguments given for register command.")
	}

	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}

	s.cfg.CurrentUserName = cmd.args[0]
	err = s.cfg.SetUser()
	if err != nil {
		return err
	}

	fmt.Printf("User %s created.\n", cmd.args[0])

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) run(s *state, cmd command) error {
	commandFunction, ok := c.commandMap[cmd.name]
	if !ok {
		return errors.New("Command not found.")
	}

	return commandFunction(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}
