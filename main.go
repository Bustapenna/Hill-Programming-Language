package main

import(
	"fmt"
	"os"
	"os/user"
	"Hill/repl"
)

func main() {
	_, err := user.Current()
	
	if err != nil {
		panic(err)
	}

	fmt.Print("Hello Creator! This is the Hill Programming Language programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
