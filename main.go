package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ath0m/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
