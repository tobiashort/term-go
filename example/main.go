package main

import (
	"fmt"

	"github.com/tobiashort/term-go"
)

func main() {
	fmt.Println("Is a tty:", term.IsTerminal())
	err := term.MakeRaw()
	if err != nil {
		panic(err)
	}
	fmt.Println("Switched to raw mode")
	err = term.Restore()
	if err != nil {
		panic(err)
	}
	fmt.Println("Restored to original mode")
}
