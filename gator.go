package main

import (
	"errors"
	"fmt"

	"github.com/joshckidd/gator/internal/config"
)

type state struct {
	config *config.Config
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

	s.config.CurrentUserName = cmd.args[0]
	err := s.config.SetUser()
	if err != nil {
		return err
	}

	fmt.Printf("User set to %s\n", cmd.args[0])
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
