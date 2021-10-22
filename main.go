package main

import (
	"fmt"
	"os"
	"os/user"
	"slang/cli"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the SLANG programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands; Use 'Ctrl + c' to exit console.\n")
	cli.Start(os.Stdin, os.Stdout)
}
