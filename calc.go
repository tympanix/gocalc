package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
	"github.com/tympanix/gocalc/scanner/token"

	"golang.org/x/crypto/ssh/terminal"
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
		term()
		os.Exit(0)
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

func term() {
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)

	t := terminal.NewTerminal(os.Stdin, "> ")
	t.AutoCompleteCallback = func(line string, pos int, key rune) (newline string, newPos int, ok bool) {
		if key == '\x03' {
			// Ctrl+C
			os.Exit(0)
		}
		return "", 0, false
	}

	for {
		text, err := t.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}

		text = strings.TrimSpace(text)

		if len(text) == 0 {
			continue
		}

		if text == "exit" || text == "quit" {
			return
		}

		s := scanner.NewFromString(text)

		p := parser.New(s).Parse()

		t.Write([]byte(fmt.Sprintln(p.Calc())))
	}
}
