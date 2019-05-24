package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
)

var (
	verbose = flag.Bool("v", false, "verbose")
)

func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalln("Missing program arguments")
	}

	s, err := scanner.NewFromFile(flag.Arg(0))

	if err != nil {
		log.Fatalln(err)
	}

	p := parser.New(s)
	n := p.ParseProgram()

	fmt.Println(n.Calc())

	if *verbose {
		n.Print()
	}
}
