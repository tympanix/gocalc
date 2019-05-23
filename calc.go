package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Missing program arguments")
	}

	s, err := scanner.NewFromFile(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	p := parser.New(s)

	n := p.ParseProgram()

	fmt.Println(n.Calc())

	n.Print()
}
