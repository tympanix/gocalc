package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
	"github.com/tympanix/gocalc/scanner/token"
)

var (
	verbose  = flag.Bool("v", false, "verbose")
	scanning = flag.Bool("s", false, "scanning")
	parsing  = flag.Bool("p", false, "parsing")
	input    = flag.String("i", "", "input")
)

func main() {

	var s *scanner.Scanner
	var err error

	flag.Parse()

	if len(*input) > 0 && flag.NArg() > 0 {
		log.Fatal("too many arguments")
	}

	if len(*input) > 0 {
		s, err = scanner.NewFromFile(*input)
	}

	if flag.NArg() > 0 {
		s = scanner.NewFromString(flag.Arg(0))
	}

	if s == nil {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			s = scanner.NewFromReader(bufio.NewReader(os.Stdin))
		}
	}

	if s == nil {
		log.Fatal("too few arguments")
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

	n.Analyze()

	if *parsing {
		n.Print()
		os.Exit(0)
	}

	fmt.Println(n.Calc())

}
