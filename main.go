package main

import (
	"github.com/mdaisuke/goscheme/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
