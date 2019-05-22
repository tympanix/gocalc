package debug

import "fmt"

var indention int

func Indent() {
	indention++
}

func Outdent() {
	indention--
}

func Println(a ...interface{}) {
	for i := 0; i < indention; i++ {
		fmt.Printf("| ")
	}
	fmt.Println(a...)
}
