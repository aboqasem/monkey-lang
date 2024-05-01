package main

import (
	"log"
	"monkey-lang/repl"
	"os"
)

func main() {
	log.Println("starting repl...")
	log.Printf("use %s to quit.", repl.QUIT)
	repl.Start(os.Stdin, os.Stdout)
}
