package main

import (
	"fmt"
	"os"
	"os/user"
	"runmark/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s.%v! This is the Monkey programming language!\n", user.Username, user.Uid)
	fmt.Printf("Feel free to type in commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
