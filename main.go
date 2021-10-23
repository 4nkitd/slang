package main

import (
	"fmt"
	"os"
	"os/user"
	"slang/cli"
)

func main() {
	_user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! \nThis is the SLANG programming language!\n", _user.Name)
	fmt.Printf("Feel free to type in commands; Use 'Ctrl + c' to exit console.\n")
	cli.Start(os.Stdin, os.Stdout)
}
