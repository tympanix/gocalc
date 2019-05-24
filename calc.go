package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
)

var (
	verbose = flag.Bool("v", false, "verbose")
	program = flag.String("p", "", "program")
)

func main() {

	var s *scanner.Scanner
	var err error

	flag.Parse()

	if len(*program) > 0 {
		r := strings.NewReader(*program)
		s, err = scanner.NewFromReader(r)
	}

	if flag.NArg() > 0 {
		s, err = scanner.NewFromFile(flag.Arg(0))
	}

	if s == nil {
		log.Fatalln("Missing program arguments")
	}
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
