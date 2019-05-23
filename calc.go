package main

import (
	"fmt"
	"log"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
)

func main() {

	s, err := scanner.New("program.p")

	if err != nil {
		log.Fatalln(err)
	}

	p := parser.New(s)

	n := p.ParseProgram()

	fmt.Println(n.Calc())

	n.Print()
}
