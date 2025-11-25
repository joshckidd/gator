package main

import (
	"fmt"
	"os"

	"github.com/joshckidd/gator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	var s state
	s.config = &c

	cliCommands := commands{
		commandMap: map[string]func(*state, command) error{},
	}

	cliCommands.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Println("No arguments provided.")
		os.Exit(1)
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}

	err = cliCommands.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
