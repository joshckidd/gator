package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joshckidd/gator/internal/config"
	"github.com/joshckidd/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", c.DBUrl)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	var s state
	s.cfg = &c
	s.db = dbQueries

	cliCommands := commands{
		commandMap: map[string]func(*state, command) error{},
	}

	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)
	cliCommands.register("users", handlerUsers)

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
