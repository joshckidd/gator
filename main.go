package main

import (
	"fmt"

	"github.com/joshckidd/gator/internal/config"
)

func main() {
	c, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	c.CurrentUserName = "josh"
	err = c.SetUser()
	if err != nil {
		fmt.Println(err)
	}

	c, err = config.Read()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
}
