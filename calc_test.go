package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/tympanix/gocalc/parser"
	"github.com/tympanix/gocalc/scanner"
)

const (
	result  = "result:"
	margin  = 1e-5
	passDir = "./test/pass"
)

func getResult(path string) (float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	r := bufio.NewScanner(f)

	for r.Scan() {
		if i := strings.Index(r.Text(), result); i > -1 {
			s := strings.TrimSpace(r.Text()[i+len(result):])
			r, err := strconv.ParseFloat(s, 64)
			if err != nil {
				return 0, err
			}
			return r, nil
		}
	}
	return 0, fmt.Errorf("missing result for file: %s", f.Name())
}

func TestPass(t *testing.T) {

	files, err := ioutil.ReadDir("./test/pass")

	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {

		t.Run(f.Name(), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Error(r)
				}
			}()

			path := path.Join(passDir, f.Name())

			s, err := scanner.NewFromFile(path)

			if err != nil {
				t.Fatal(err)
			}

			res, err := getResult(path)

			if err != nil {
				t.Fatal(err)
			}

			n := parser.New(s).Parse()
			n.Analyze()
			r := n.Calc()

			if r > res+margin || r < res-margin {
				t.Errorf("result: %f, expected: %f", r, res)
			}

		})

	}

}
