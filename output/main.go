package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! This is the Monkey programming language!", u.Username)
	fmt.Println("Fell free to type in commands")

	repl.Start(os.Stdin, os.Stdout)
}
