package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
	"github.com/tympanix/gocalc/scanner/token"
)

var (
	verbose  = flag.Bool("v", false, "verbose")
	scanning = flag.Bool("s", false, "scanning")
	program  = flag.String("p", "", "program")
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

	if *scanning {
		for {
			t := s.NextToken()
			if t.Kind() == token.EOF {
				return
			}
			fmt.Printf("%-12s: %s\n", t.Kind().String(), t.String())
		}
	}

	n := parser.New(s).Parse()

	fmt.Println(n.Calc())

	if *verbose {
		n.Print()
	}
}
