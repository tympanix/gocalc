package main

import (
	"fmt"
	"log"

	"github.com/tympanix/tymlang/parser"
	"github.com/tympanix/tymlang/scanner"
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
