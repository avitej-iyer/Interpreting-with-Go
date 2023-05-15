package main

import (
	"fmt"
	"Interpreting-with-Go/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is my language Splice (currently under developement)\n",
		user.Username)
	fmt.Printf("Please type in commands below\n")
	repl.Start(os.Stdin, os.Stdout)
}
