package main

import (
	"time"

	"github.com/tobiashort/term-go"
)

func main() {
	err := term.MakeRaw()
	if err != nil {
		panic(err)
	}
	time.Sleep(2 * time.Second)
	err = term.Restore()
	if err != nil {
		panic(err)
	}
}
